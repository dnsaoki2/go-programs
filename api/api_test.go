package main 

import "testing"

func TestUnMarshal(t *testing.T) {
	input := []byte(`{"subtitulo":"subtitulo","foto":"foto","chapeu":"chapeu","url":"url","estilo":"estilo","titulo":"titulo"}`)
	result := unMarshal(input)
	if result.Subtitulo != "subtitulo" {
		t.Error("Expected subtitulo, got", result.Subtitulo)
	}
	if result.Foto != "foto" {
		t.Error("Expected foto, got", result.Foto)
	}
	if result.Chapeu != "chapeu" {
		t.Error("Expected chapeu, got", result.Chapeu)
	}
	if result.Url != "url" {
		t.Error("Expected url, got", result.Url)
	}
	if result.Estilo != "estilo" {
		t.Error("Expected estilo, got", result.Estilo)
	}
	if result.Titulo != "titulo" {
		t.Error("Expected titulo, got", result.Titulo)
	}
}


func TestSplit(t *testing.T) {
	input := []byte(`[{"subtitulo":"subtitulo","foto":"foto","chapeu":"chapeu","url":"url","estilo":"estilo","titulo":"titulo"},{"subtitulo":"subtitulo1","foto":"foto1","chapeu":"chapeu1","url":"url1","estilo":"estilo1","titulo":"titulo1"}]`)
	result := split(input)
	resultMar0 := unMarshal([]byte(result[0]))
	resultMar1 := unMarshal([]byte(result[1]))
	if resultMar0.Subtitulo != "subtitulo" {
		t.Error("Expected subtitulo, got ", resultMar0.Subtitulo)
	}
	if resultMar1.Subtitulo != "subtitulo1" {
		t.Error("Expected subtitulo1, got ", resultMar1.Subtitulo)
	}
	if resultMar0.Foto != "foto" {
		t.Error("Expected foto, got ", resultMar0.Foto)
	}
	if resultMar1.Foto != "foto1" {
		t.Error("Expected foto1, got ", resultMar1.Foto)
	}
	if resultMar0.Chapeu != "chapeu" {
		t.Error("Expected chapeu, got ", resultMar0.Chapeu)
	}
	if resultMar1.Chapeu != "chapeu1" {
		t.Error("Expected chapeu1, got ", resultMar1.Chapeu)
	}
	if resultMar0.Url != "url" {
		t.Error("Expected url, got ", resultMar0.Url)
	}
	if resultMar1.Url != "url1" {
		t.Error("Expected url1, got ", resultMar1.Url)
	}
	if resultMar0.Estilo != "estilo" {
		t.Error("Expected estilo, got ", resultMar0.Estilo)
	}
	if resultMar1.Estilo != "estilo1" {
		t.Error("Expected estilo1, got ", resultMar1.Estilo)
	}
	if resultMar0.Titulo != "titulo" {
		t.Error("Expected titulo, got ", resultMar0.Titulo)
	}
	if resultMar1.Titulo != "titulo1" {
		t.Error("Expected titulo1, got ", resultMar1.Titulo)
	}
}
