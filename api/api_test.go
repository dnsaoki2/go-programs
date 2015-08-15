package main 

import (
	"testing"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

var _ = check.Suite(RecSuite{})

type RecSuite struct{}


func (RecSuite) TestRead(c *check.C) {
	input := []byte('{"subtitulo":"subtitulo","foto":"foto","chapeu":"chapeu","url":"url","estilo":"estilo","titulo":"titulo"}')
	result := read(input)
}
