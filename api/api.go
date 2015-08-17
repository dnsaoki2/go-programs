package main 

import (
  "encoding/json"
  "io/ioutil"
  "html/template"
  "net/http"
  "strings"
  "fmt"
 )

type Page struct {
  Subtitulo   string
  Foto		    string
  Chapeu		  string
  Url			    string
  Estilo		  string
  Titulo		  string
}

//Split the file into multiple objects
func split(file []byte) []string {
  str := string(file)
  str = strings.Trim(str, "[")
  str = strings.Trim(str,"]")
  str = strings.TrimSpace(str)
  strSplit := strings.Split(str, "},")
  for i := 0; i < len(strSplit); i++ {
    strAux := make([]byte, len(strSplit[i]))
    if !strings.Contains(strSplit[i],"}") {
      strAux = append([]byte(strSplit[i]),'}')
      strSplit[i] = string(strAux)
    }
  } 
  return strSplit
}
  
//Unpack the objects
func unMarshal(file []byte) Page {
	var data Page
  json.Unmarshal(file, &data)
 	return data
}

//start server
func startServer() {
  http.HandleFunc("/news", page)
  http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("templates/CSS"))))
  http.ListenAndServe(":8080", nil)
}

//Template with data 
func page(w http.ResponseWriter, r *http.Request) {
  uf := r.URL.Query().Get("uf")
  if uf == "" {
    http.Error(w, "UF invalido", http.StatusInternalServerError)
    return
  }
  site := fmt.Sprintf("http://c.api.globo.com/news/%s.json", uf)
  file, err := http.Get(site)
  if err != nil {
    panic(err)
  }
  dataByte, err := ioutil.ReadAll(file.Body)
  if err != nil {
    panic(err)
  }
  dataSplit := split(dataByte)
  for i := 0; i < len(dataSplit); i++ { 
    dataPage := unMarshal([]byte(dataSplit[i]))
    tmpl, err := template.ParseFiles("./templates/index.html")
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    if err := tmpl.Execute(w, dataPage); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}

func main() {
  startServer()
}

