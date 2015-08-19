package main

import (
	"gopkg.in/check.v1"
	"testing"
)

var unMarshalInput = []byte(`{"subtitulo":"subtitulo","foto":"foto","chapeu":"chapeu","url":"url","estilo":"estilo","titulo":"titulo"}`)
var splitInput = []byte(`[{"subtitulo":"subtitulo","foto":"foto","chapeu":"chapeu","url":"url","estilo":"estilo","titulo":"titulo"},{"subtitulo":"subtitulo1","foto":"foto1","chapeu":"chapeu1","url":"url1","estilo":"estilo1","titulo":"titulo1"}]`)

func Test(t *testing.T) {
	check.TestingT(t)
}

type testingApi struct{}

var _ = check.Suite(&testingApi{})

func (t *testingApi) TestUnMarshal(c *check.C) {
	result := unMarshal(unMarshalInput)
	c.Check(result.Subtitulo, check.Equals, "subtitulo")
	c.Check(result.Foto, check.Equals, "foto")
	c.Check(result.Chapeu, check.Equals, "chapeu")
	c.Check(result.Url, check.Equals, "url")
	c.Check(result.Estilo, check.Equals, "estilo")
	c.Check(result.Titulo, check.Equals, "titulo")
}

func (t *testingApi) TestSplit(c *check.C) {
	result := split(splitInput)
	resultMar0 := unMarshal([]byte(result[0]))
	resultMar1 := unMarshal([]byte(result[1]))
	c.Check(resultMar0.Subtitulo, check.Equals, "subtitulo")
	c.Check(resultMar1.Subtitulo, check.Equals, "subtitulo1")
	c.Check(resultMar0.Foto, check.Equals, "foto")
	c.Check(resultMar1.Foto, check.Equals, "foto1")
	c.Check(resultMar0.Chapeu, check.Equals, "chapeu")
	c.Check(resultMar1.Chapeu, check.Equals, "chapeu1")
	c.Check(resultMar0.Url, check.Equals, "url")
	c.Check(resultMar1.Url, check.Equals, "url1")
	c.Check(resultMar0.Estilo, check.Equals, "estilo")
	c.Check(resultMar1.Estilo, check.Equals, "estilo1")
	c.Check(resultMar0.Titulo, check.Equals, "titulo")
	c.Check(resultMar1.Titulo, check.Equals, "titulo1")
}
