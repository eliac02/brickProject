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
		temp := strings.Split(scanner.Text(), " ")
		switch temp[0] {
		case "m":
			inserisciMattoncino(g, temp[1], temp[2], temp[3])
		case "s":
			stampaMattoncino(g, temp[1])
			fmt.Println()
		case "d":
			listaNomi := ""
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
