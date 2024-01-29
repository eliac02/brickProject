package main

import (
	"bufio"
	"fmt"
	"os"
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

/*type fila struct {
	componenti      *linkedList
	nome            string
	indiceCacofonia int
}*/

func main() {
	var g gioco = gioco{make(map[string][2]string), make(map[string]*linkedList)}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		operation, operand, _ := strings.Cut(scanner.Text(), " ")
		switch operation {
		case "m":
			names := strings.Split(operand, " ")
			inserisciMattoncino(g, names[0], names[1], names[2])
		case "s":
			stampaMattoncino(g, operand)
			fmt.Println()
		case "d":
			listaNomi := operand //per esempio "+ciao -cane -gatto +macchina"
			disponiFila(g, listaNomi)
		case "S":
			stampaFila(g, temp[1])
		case "e":
			eliminaFila(g, temp[1])
		case "q":
			return
		}
	}
}
