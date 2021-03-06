package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/pat"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
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

type UfPage struct {
	Uf     string
	Pageuf []Page
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
	r := pat.New()
	r.Get("/news/{uf:[A-Za-z]+}", requestPageDB)
	http.Handle("/", r)
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("templates/CSS"))))
	http.ListenAndServe(":2020", nil)
}

//func to connect database
func connectionDB() *mgo.Session {
	//session, err := mgo.Dial(os.Getenv("DB_PORT_27017_TCP_ADDR")+":"+os.Getenv("DB_PORT_27017_TCP_PORT"))
	session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
	if err != nil {
		panic(err)
	}
	return session
}

//Request from users
func requestPageDB(w http.ResponseWriter, r *http.Request) {
	uf := strings.ToUpper(r.URL.Query().Get(":uf"))
	if !contains(uf) {
		http.Error(w, "Invalid UF", http.StatusInternalServerError)
		return
	}
	session := connectionDB()
	defer session.Close()
	var result []UfPage
	err := session.DB("apidb").C("news").Find(bson.M{"uf": uf}).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	for index := range result[0].Pageuf {
		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, result[0].Pageuf[index])
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
	var wait sync.WaitGroup
	fmt.Println("Update start ...")
	for index := range ufs {
		wait.Add(1)
		copy := index
		go func() {
			savePageMemory_(copy)
			wait.Done()
		}()
	}
	wait.Wait()
	fmt.Println(" ...Update done")
}

//goroutine to save page ufs[index]
func savePageMemory_(index int) {
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
	session := connectionDB()
	defer session.Close()
	c := session.DB("apidb").C("news")
	var ufPage UfPage
	ufPage.Uf = ufs[index]
	ufPage.Pageuf = make([]Page, len(dataSplit))
	for i := 0; i < len(dataSplit); i++ {
		ufPage.Pageuf[i] = unMarshal([]byte(dataSplit[i]))
	}
	_, err = c.Upsert(bson.M{"uf": ufs[index]}, &ufPage)
	if err != nil {
		log.Fatal(err)
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
	time.Sleep(10 * time.Second)
	startServer()
}
