package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type gioco struct {
	scatola map[string][2]string //la chiave e' il nome del mattoncino, il valore sono i due bordi
	tavolo  map[string]fila      //la chiave e' il nome della fila, il valore e' la fila
}

type mattoncino struct {
	alpha string
	beta  string
	sigma string
}

type fila struct {
	componenti      *linkedList
	indiceCacofonia int
}

func main() {
	var g gioco = gioco{make(map[string][2]string), make(map[string]fila)}

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
			stampaFila(g, operand)
		case "e":
			eliminaFila(g, operand)
		case "f":
			estremi := strings.Split(operand, " ")
			disponiFilaMinima(g, estremi[0], estremi[1])
		case "i":
			fmt.Println(indiceCacofonia(g, operand))
		case "M":
			stringhe := strings.Split(operand, " ")
			fmt.Println(sottostringaMassima(stringhe[0], stringhe[1]))
		case "q":
			return
		}
	}
}
