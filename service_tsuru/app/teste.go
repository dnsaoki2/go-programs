package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	//"os"
)

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
	Pageuf Page
}

func main() {
	//host := os.Getenv("MONGODB_URI")
	//host := "127.0.0.1:27017"
	// host := "192.168.99.100:27017"
	// session, err := mgo.Dial(host)
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()
	// c := session.DB("apidb").C("news")
	// var ufPage UfPage
	// ufPage.Uf = "teste"
	// ufPage.Pageuf = Page{"sub", "foto", "chapeu", "url", "estilo", "titulo"}
	// _, err = c.Upsert(bson.M{"uf": "teste"}, &ufPage)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var result []UfPage
	// err = session.DB("apidb").C("news").Find(bson.M{"uf": "teste"}).All(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("teste2")
	fmt.Println("teste1")
}
