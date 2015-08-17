package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "http://c.api.globo.com/news/RJ.json"
  re := regexp.MustCompile(`http://c.api.globo.com/([A-Za-z/]+).json$`)
  e := re.FindStringSubmatch(s)
  fmt.Println(e)
}
