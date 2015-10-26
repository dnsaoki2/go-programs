package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
	fmt.Println(strings.HasSuffix("regiv3ry.tsuru.globoi.com/tsuru/app-example:v2", ":v2"))
	fmt.Println(strings.HasSuffix("regiv3ry.t:v4uru.globoi.com/tsuru/app-example:v2", ":v"))
}
