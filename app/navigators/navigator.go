package navigators

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/xmlquery"
	"github.com/chromedp/chromedp"
)

func ObtemDetalheVoo(term string) (DetalheVoo, error) {
	detalhe, err := ObtemInformacoesDetalheVoo(term)
	if err != nil {
		return DetalheVoo{}, err
	}

	b, err := json.Marshal(detalhe)
	if err != nil {
		fmt.Println(err)
		return DetalheVoo{}, err
	}
	downloadUnzip(term, detalhe.LinkFile)
	files, err := ioutil.ReadDir(term)
	if err != nil {
		return DetalheVoo{}, err
	}

	fileUnziped := fmt.Sprintf("%s/%s", term, files[0].Name())
	content, err := ioutil.ReadFile(fileUnziped)

	if err != nil {
		return DetalheVoo{}, err
	}

	track := readXMLTrackLog(string(content))
	detalhe.TrackingLog = track
	b, err = json.Marshal(detalhe)
	if err != nil {
		fmt.Println(err)
		return DetalheVoo{}, err
	}
	e := os.Remove(fileUnziped)
	e = os.Remove(term)
	e = os.Remove(fmt.Sprintf("%s.zip", term))

	if e != nil {
		return DetalheVoo{}, err
	}
	f, err := os.Create(fmt.Sprintf("%s.json", term))

	if err != nil {
		return DetalheVoo{}, err
	}

	defer f.Close()

	_, err2 := f.Write(b)

	if err2 != nil {
		return DetalheVoo{}, err
	}

	return detalhe, nil
}

func downloadUnzip(name string, link string) {
	downloadPackage(name, link)
	unzip(fmt.Sprintf("%s.zip", name), name)
}

func downloadPackage(name string, link string) {
	url := fmt.Sprintf(link)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	filename := fmt.Sprintf("%s.zip", name)
	out, _ := os.Create(filename)
	defer out.Close()
	io.Copy(out, resp.Body)
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

type Voo struct {
	Ordem              string `json:"ordem"`
	Data               string `json:"data"`
	Piloto             string `json:"piloto"`
	HoraDecolagem      string `json:"horaDecolagem"`
	Decolagem          string `json:"decolagem"`
	HoraPouso          string `json:"horaPouso"`
	Pouso              string `json:"pouso"`
	Duracao            string `json:"duracao"`
	DistanciaLinhaReta string `json:"distanciaLinhaReta"`
	OLCKm              string `json:"olcKm"`
	PontuacaoOLC       string `json:"pontuacaoOlc"`
	Link               string `json:"link"`
	Equipamento        string `json:"equipamento"`
	MaiorAscendente    string `json:"maiorAscendente"`
	MaiorDescendente   string `json:"maiordescendente"`
	AltitudeMaxima     string `json:"altitudeMaxima"`
	AltitudeMinima     string `json:"altitudeMinima"`
	AltitudeDecolagem  string `json:"altitudeDecolagem"`
	GanhoAltitude      string `json:"ganhoAltitude"`
	VelocidadeMaxima   string `json:"velocidadeMaxima"`
	VelocidadeMedia    string `json:"velocidadeMedia"`
}

type DetalheVoo struct {
	Voo         Voo         `json:"voo"`
	TrackingLog TrackingLog `json:"trackingLog"`
	LinkFile    string      `json:"linkFile"`
}
type TrackingLog struct {
	Nome        string   `json:"nome"`
	DisplayName string   `json:"displayName"`
	Times       []string `json:"times"`
	Coords      []Coord  `json:"coords"`
	Climbs      []string `json:"climbs"`
}
type Coord struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
	Alt float64 `json:"altitude"`
}

func ObtemListaVoos() ([]Voo, error) {
	listaVoos, err := GetHttpHtmlContent("http://xcbrasil.com.br/", "body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table", `document.querySelector("body")`)
	if err != nil {
		panic(err)
	}
	lista, err := ParseListaVoos(listaVoos)
	if err != nil {
		return []Voo{}, err
	}
	return lista, nil
}

func ObtemInformacoesDetalheVoo(id string) (DetalheVoo, error) {
	site := fmt.Sprintf(`http://xcbrasil.com.br/flight/%s`, id)
	detalheVoo, err := GetHttpHtmlContent(site, "#pilot_pos", `document.querySelector("body")`)
	if err != nil {
		panic(err)
	}
	detalhe, err := ParseDetalheVoo(detalheVoo)
	if err != nil {
		return DetalheVoo{}, err
	}
	return detalhe, nil
}

func ParseDetalheVoo(conteudo string) (DetalheVoo, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(conteudo))
	if err != nil {
		log.Fatal(err)
		return DetalheVoo{}, err
	}
	piloto := dom.Find("#pilot_pos > a").First().Text()
	data := strings.Split(dom.Find("#pilot_pos").First().Text(), "Data: ")[1]
	horaDecolagem := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(1) > table:nth-child(1) > tbody > tr:nth-child(2) > td:nth-child(2) > div > span").First().Text()
	decolagem := dom.Find("#takeoffAddPos > a").First().Text()
	horaPouso := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(1) > table:nth-child(1) > tbody > tr:nth-child(4) > td:nth-child(2) > div > span").First().Text()
	pouso := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(1) > table:nth-child(1) > tbody > tr:nth-child(5) > td > div").First().Text()
	duracao := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(2) > table:nth-child(2) > tbody > tr:nth-child(2) > td:nth-child(2) > div > span").First().Text()
	distanciaLinhaReta := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(1) > table:nth-child(2) > tbody > tr:nth-child(3) > td:nth-child(2) > div").First().Text()
	olcKm := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(1) > table:nth-child(2) > tbody > tr:nth-child(4) > td:nth-child(2) > div").First().Text()
	pontuacaoOlc := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(1) > table:nth-child(2) > tbody > tr:nth-child(5) > td:nth-child(2) > div > strong").First().Text()
	equipamento := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(2) > table:nth-child(1) > tbody > tr:nth-child(2) > td > div > div").First().Text()
	maiorAscendente := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(2) > table:nth-child(2) > tbody > tr:nth-child(3) > td:nth-child(2) > div > span").First().Text()
	maiorDescendente := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(2) > table:nth-child(2) > tbody > tr:nth-child(4) > td:nth-child(2) > div > span").First().Text()
	altitudeMaxima := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(2) > table:nth-child(2) > tbody > tr:nth-child(5) > td:nth-child(2) > div > span").First().Text()
	altitudeMinima := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(2) > table:nth-child(2) > tbody > tr:nth-child(6) > td:nth-child(2) > div > span").First().Text()
	altitudeDecolagem := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(2) > table:nth-child(2) > tbody > tr:nth-child(7) > td:nth-child(2) > div > span").First().Text()
	ganhoAltitude := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(2) > table:nth-child(2) > tbody > tr:nth-child(8) > td:nth-child(2) > div > span").First().Text()
	velocidadeMaxima := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(2) > table:nth-child(2) > tbody > tr:nth-child(9) > td:nth-child(2) > div > span").First().Text()
	velocidadeMedia := dom.Find("body > table > tbody > tr > td:nth-child(1) > table:nth-child(5) > tbody > tr > td.mainWrapCell > div.main_text > table > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(2) > table:nth-child(2) > tbody > tr:nth-child(10) > td:nth-child(2) > div > span").First().Text()
	voo := Voo{
		Piloto:             piloto,
		Data:               data,
		HoraDecolagem:      horaDecolagem,
		Decolagem:          decolagem,
		HoraPouso:          horaPouso,
		Pouso:              pouso,
		Duracao:            duracao,
		DistanciaLinhaReta: distanciaLinhaReta,
		OLCKm:              olcKm,
		PontuacaoOLC:       pontuacaoOlc,
		Equipamento:        equipamento,
		MaiorAscendente:    maiorAscendente,
		MaiorDescendente:   maiorDescendente,
		AltitudeMaxima:     altitudeMaxima,
		AltitudeMinima:     altitudeMinima,
		AltitudeDecolagem:  altitudeDecolagem,
		GanhoAltitude:      ganhoAltitude,
		VelocidadeMaxima:   velocidadeMaxima,
		VelocidadeMedia:    velocidadeMedia,
	}

	partLink, _ := dom.Find("#ge_s0").First().Attr("href")
	linkfile := fmt.Sprintf(`http://xcbrasil.com.br%s`, partLink)
	resp, _ := http.Get(linkfile)
	defer resp.Body.Close()
	xml, _ := io.ReadAll(resp.Body)
	link := readXMLVoo(string(xml))

	detalhe := DetalheVoo{
		LinkFile: link,
		Voo:      voo,
	}
	return detalhe, nil
}

func readXMLVoo(s string) string {
	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	href := xmlquery.FindOne(doc, "//NetworkLink/Link/href")
	hrefText := href.InnerText()
	//hrefText = hrefText[1 : len(hrefText)-1]
	return hrefText
}

func readXMLTrackLog(s string) TrackingLog {

	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	track := TrackingLog{}
	root := xmlquery.FindOne(doc, "//kml/Document")
	if n := root.SelectElement("name"); n != nil {
		track.Nome = n.InnerText()
	}
	if n := root.SelectElement("Schema/gx:SimpleArrayField/displayName"); n != nil {
		track.DisplayName = n.InnerText()
	}
	var times []string
	for _, t := range xmlquery.Find(doc, "//kml/Document/Placemark[1]/gx:MultiTrack[1]/gx:Track/when") {
		//fmt.Printf("Times #%d %s\n", i, times.InnerText())
		times = append(times, t.InnerText())

	}
	var coords []Coord
	for _, c := range xmlquery.Find(doc, "//kml/Document/Placemark[1]/gx:MultiTrack[1]/gx:Track/gx:coord") {
		coord := strings.Split(c.InnerText(), " ")
		lng, _ := strconv.ParseFloat(coord[0], 8)
		lat, _ := strconv.ParseFloat(coord[1], 8)
		alt, _ := strconv.ParseFloat(coord[2], 8)
		coords = append(coords, Coord{
			Lng: lng,
			Lat: lat,
			Alt: alt,
		})

	}
	var climbs []string
	for _, cl := range xmlquery.Find(doc, "//kml/Document/Placemark[1]/gx:MultiTrack[1]/gx:Track/ExtendedData/SchemaData/gx:SimpleArrayData/gx:value") {
		climbs = append(climbs, cl.InnerText())
	}
	track.Times = times
	track.Coords = coords
	track.Climbs = climbs
	return track
}

func ParseListaVoos(conteudo string) ([]Voo, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(conteudo))
	if err != nil {
		log.Fatal(err)
		return []Voo{}, err
	}
	var lista []Voo
	elemento := dom.Find(".listTable").First()
	elemento.Find("tr").Each(func(i int, tr *goquery.Selection) {
		if i > 0 {
			link, _ := tr.Find(".flightLink").Attr("href")
			voo := Voo{
				Ordem:              tr.Find(".indexCell").Text(),
				Data:               tr.Find(".dateString").Text(),
				Piloto:             tr.Find(".pilotLink").Text(),
				Decolagem:          tr.Find(".takeoffLink").Text(),
				Duracao:            tr.Find("td:nth-child(4)").Text(),
				DistanciaLinhaReta: tr.Find("td:nth-child(5)").Text(),
				OLCKm:              tr.Find("td:nth-child(6)").Text(),
				PontuacaoOLC:       tr.Find("td:nth-child(7)").Text(),
				Link:               link,
			}
			lista = append(lista, voo)
		}
	})
	return lista, nil
}

func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug using
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	//Initialization parameters, first pass an empty data
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	//Execute an empty task to create a chrome instance in advance
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//Create a context with a timeout of 40s
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		return "", err
	}
	return htmlContent, nil
}
