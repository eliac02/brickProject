package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// gioco represents the game state.
type gioco struct {
	scatola map[string][2]string //la chiave e' il nome del mattoncino, il valore sono i due bordi
	tavolo  map[string]fila      //la chiave e' il nome della fila, il valore e' la fila
	forme   map[string]bool      //la chiave e' la forma, il valore e' booleano
}

// mattoncino represents a brick with specific shapes and name.
type mattoncino struct {
	alpha string
	beta  string
	sigma string
}

// fila represents a row in the game.
type fila struct {
	componenti      *linkedList
	indiceCacofonia int
}

// This program reads input from stdin and executes operations based on the specified command.
//
// Commands:
// - "m": Insert a new brick.
// - "s": Print information about a brick.
// - "d": Arrange a row.
// - "S": Print a row.
// - "e": Delete a row.
// - "f": Arrange a row in the shortest way.
// - "i": Calculate the cacophony index of a row.
// - "M": Find the maximum substring between two strings.
// - "q": Quit the program.
func main() {
	var g gioco = gioco{make(map[string][2]string), make(map[string]fila), make(map[string]bool)}

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
		case "c":
			nome, forme, _ := strings.Cut(operand, " ")
			costo(g, nome, forme)
		case "q":
			return
		}
	}
}
