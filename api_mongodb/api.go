package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//global variables
var ufs = []string{"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG",
	"PR", "PB", "PA", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SE", "SP", "TO"}

type Page struct {
	Subtitulo string
	Foto      string
	Chapeu    string
	Url       string
	Estilo    string
	Titulo    string
}

//Split the file into multiple objects
func split(file []byte) []string {
	str := string(file)
	str = strings.Trim(str, "[")
	str = strings.Trim(str, "]")
	strSplit := strings.Split(str, "},")
	for i := 0; i < len(strSplit); i++ {
		strAux := make([]byte, len(strSplit[i]))
		if !strings.Contains(strSplit[i], "}") {
			strAux = append([]byte(strSplit[i]), '}')
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
	http.HandleFunc("/news", requestPageDB)
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("templates/CSS"))))
	http.ListenAndServe(":8080", nil)
}

func connectionDB() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	return session
}

//Request from users
func requestPageDB(w http.ResponseWriter, r *http.Request) {
	uf := strings.ToUpper(r.URL.Query().Get("uf"))
	if !contains(uf) {
		http.Error(w, "Invalid UF", http.StatusInternalServerError)
		return
	}
	session := connectionDB()
	defer session.Close()
	var result []Page
	err := session.DB("apidb").C(uf).Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	for index := range result {
		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, result[index])
		if err != nil {
			panic(err)
		}
	}
}

//Func to update the web pages in memory
func upd() {
	savePageMemory()
	for range time.Tick(time.Minute) {
		savePageMemory()
	}
}

//Save the web page in memory
func savePageMemory() {
	fmt.Println("Update memory")
	for index := range ufs {
		copy := index
		go func() {
			savePageMemory_(copy)
		}()
	}
}

//goroutine to save page ufs[index]
func savePageMemory_(index int) {
	//Get Site for ufs[index]
	site := fmt.Sprintf("http://c.api.globo.com/news/%s.json", ufs[index])
	file, err := http.Get(site)
	//connection error
	if err != nil {
		panic(err)
	}
	//get all data from website
	dataByte, err := ioutil.ReadAll(file.Body)
	if err != nil {
		panic(err)
	}
	//split the data
	dataSplit := split(dataByte)
	for i := 0; i < len(dataSplit); i++ {
		//Unpack the objects
		dataPage := unMarshal([]byte(dataSplit[i]))
		//Connection to database and save the data struct in db
		session := connectionDB()
		defer session.Close()
		c := session.DB("apidb").C(ufs[index])
		err = c.Insert(&dataPage)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//Func to verify if ufs contais string s
func contains(s string) bool {
	for _, value := range ufs {
		if strings.EqualFold(s, value) {
			return true
		}
	}
	return false
}

func main() {
	go upd()
	time.Sleep(2 * time.Second)
	startServer()

}
