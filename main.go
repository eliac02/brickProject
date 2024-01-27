package main

import (
	"fmt"
)

type gioco struct {
	scatola map[string][2]string
	tavolo  []fila
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
	g := gioco{scatola: make(map[string][2]string), tavolo: make([]fila, 0)}
	listaNomi := ""
	disponiFila(g, listaNomi)
}
