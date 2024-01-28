package main

import (
	"fmt"
	"strings"
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
	componenti      *linkedList
	nome            string
	indiceCacofonia int
}

func main() {
	var istruzione string
	var g gioco = gioco{make(map[string][2]string), make(map[string]*linkedList)}
	for {
		fmt.Scan(&istruzione)
		temp := strings.Split(istruzione, " ")
		switch temp[0] {
		case "m":
			inserisciMattoncino(g, temp[1], temp[2], temp[3])
		case "s":
			stampaMattoncino(g, temp[1])
		case "d":
			listaNomi := ""
			disponiFila(g, listaNomi)
		case "S":
			stampaFila(g, temp[1])
		case "e":
			eliminaFila(g, temp[1])
		case "q":
			break
		}
	}
}
