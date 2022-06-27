package navigators

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

type Voo struct {
	Ordem              string `json:"ordem"`
	Data               string `json:"data"`
	Piloto             string `json:"piloto"`
	Decolagem          string `json:"decolagem"`
	Duracao            string `json:"duracao"`
	DistanciaLinhaReta string `json:"distanciaLinhaReta"`
	OLCKm              string `json:"olcKm"`
	PontuacaoOLC       string `json:"pontuacaoOlc"`
	Link               string `json:"link"`
}

func ListaVoos() ([]Voo, error) {
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
func DetalheVoo(id string) (Voo, error) {
	site := fmt.Sprintf(`http://xcbrasil.com.br/flight/%s`, id)
	log.Println("site:", site)
	detalheVoo, err := GetHttpHtmlContent(site, "#pilot_pos", `document.querySelector("body")`)
	if err != nil {
		panic(err)
	}
	detalhe, err := ParseDetalheVoo(detalheVoo)
	if err != nil {
		return Voo{}, err
	}
	return detalhe, nil
}
func ParseDetalheVoo(conteudo string) (Voo, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(conteudo))
	if err != nil {
		log.Fatal(err)
		return Voo{}, err
	}
	var detalhe Voo
	partLink, _ := dom.Find("#ge_s0").First().Attr("href")
	linkfile := fmt.Sprintf(`http://xcbrasil.com.br/%s`, partLink)
	xmlvoo := downloadUnzip(linkfile)
	detalhe = Voo{Ordem: "teste", Piloto: xmlvoo}
	return detalhe, nil

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

func downloadUnzip(link string) string {
	log.Println("nome arquivo:", link)
	download(link, "temporario")
	unzip("temporario.zip", "temporario")
	files, err := ioutil.ReadDir("temporario")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
	errRemo := os.Remove("temporario*")
	if errRemo != nil {
		log.Fatal(err)
	}
	return ""
}

func download(name string, link string) {
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
