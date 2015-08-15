package main 

import (
  "encoding/json"
  "io/ioutil"
  "html/template"
  "net/http"
 )

type Pagina struct {
  Subtitulo 	string
  Foto		    string
  Chapeu		  string
  Url			    string
  Estilo		  string
  Titulo		  string
}

func read(file []byte) Pagina {
	var data Pagina
  json.Unmarshal(file, &data)
 	return data
}

func initServer() {
	http.HandleFunc("/", pag)
  http.ListenAndServe(":3000", nil)
}

func pag(w http.ResponseWriter, r *http.Request) {
  file, e := ioutil.ReadFile("list.json")
  if e != nil {
    panic(e)
  }
  dataPagina := read(file) 
  tmpl, err := template.ParseFiles("./templates/index.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  if err := tmpl.Execute(w, dataPagina); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func main() {
    initServer()
}

