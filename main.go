package main

import (
	"fmt"
)

type gioco struct {
	scatola map[string][2]string //la chiave e' il nome del mattoncino, il valore sono i due bordi
	tavolo  map[string]*linkedList
}

type mattoncino struct {
	alpha string
	beta  string
	sigma string
}

type fila struct {
	componenti      linkedList
	nome            string
	indiceCacofonia int
}

func main() {
	fmt.Println("Hello, playground")
	g := gioco{scatola: make(map[string][2]string), tavolo: make(map[string]*linkedList)}
	listaNomi := ""
	disponiFila(g, listaNomi)
}
