//teste
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	//"os"
	"strings"
	"sync"
	"time"
)

//global variables
var memory map[string]*bytes.Buffer
var lock sync.Mutex
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
	http.HandleFunc("/news", requestPageBuffer)
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("templates/CSS"))))
	http.ListenAndServe(":8080", nil)
}

//Request from users
func requestPageBuffer(w http.ResponseWriter, r *http.Request) {
	uf := strings.ToUpper(r.URL.Query().Get("uf"))
	if !contains(uf) {
		http.Error(w, "Invalid UF", http.StatusInternalServerError)
		return
	}
	lock.Lock()
	temp := *memory[uf]
	memory[uf].WriteTo(w)
	memory[uf] = &temp
	lock.Unlock()
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
	memoryTmp := make(map[string]*bytes.Buffer)
	var wait sync.WaitGroup
	for index := range ufs {
		wait.Add(1)
		copy := index
		go func() {
			memoryTmp[ufs[copy]] = savePageMemory_(copy)
			wait.Done()
		}()
	}
	wait.Wait()
	lock.Lock()
	memory = memoryTmp
	lock.Unlock()
}

//goroutine to save page ufs[index]
func savePageMemory_(index int) *bytes.Buffer {
	memoryTmp := make(map[string]*bytes.Buffer)
	bufferTmp := new(bytes.Buffer)
	//Get Site for ufs[index]
	site := fmt.Sprintf("http://c.api.globo.com/news/%s.json", ufs[index])
	file, err := http.Get(site)
	//connection error
	if err != nil {
		panic(err)
	}
	//status error
	if file.StatusCode != 200 {
		bufferTmp = bytes.NewBuffer([]byte(string(file.StatusCode)))
	} else {
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
			//make template
			tmpl, err := template.ParseFiles("./templates/index.html")
			if err != nil {
				panic(err)
			}
			//save in a buffer
			err = tmpl.Execute(bufferTmp, dataPage)
			if err != nil {
				panic(err)
			}
		}
	}
	memoryTmp[ufs[index]] = bufferTmp
	return memoryTmp[ufs[index]]
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
	startServer()
}
