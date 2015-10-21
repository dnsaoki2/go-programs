package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ins struct {
	uf    string
	teste string
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017") //os.Getenv("MONGODB_URI"))
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("apidb").C("news")
	str := ins{"1", "t"}
	err = c.Insert(str)
	if err != nil {
		log.Fatal(err)
	}
	tam, err := c.Find(bson.M{"teste": "t"}).Count()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(tam)
}
