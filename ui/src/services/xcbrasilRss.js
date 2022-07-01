import axios from "axios";
const baseURL = `http://www.xcbrasil.com.br/rss.php?op=latest&pilot=`
const getPiloto = (piloto) => axios.get(baseURL+piloto)
