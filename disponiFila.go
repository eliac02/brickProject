package main

import "strings"

// creaFila creates a row in the game.
//
// @param g The game object.
// @param sliceNames The names of the bricks that will form the row.
// @param list The linked list that represents the row.
func creaFila(g gioco, sliceNames []string, list *linkedList) {
	for _, name := range sliceNames {
		node := newNode(mattoncino{alpha: g.scatola[name[1:]][0], beta: g.scatola[name[1:]][1], sigma: name[1:]}, name[0])
		addNode(list, node)
		delete(g.scatola, name[1:])
		delete(g.forme, name[1:])
	}
}

// disponiFila arranges a row of bricks on the table.
//
// @param g The game object.
// @param listaNomi The string containing the names of the bricks that will form the row.
func disponiFila(g gioco, listaNomi string) {
	var nomeFila string
	sliceNames := strings.Split(listaNomi, " ")

	for _, name := range sliceNames {
		temp := name[1:]
		if !controllaScatola(g.scatola, temp) {
			return
		}
		nomeFila += temp + " "
	}

	if verificaFila(sliceNames, g) {
		l := newList()
		creaFila(g, sliceNames, l)
		g.tavolo[nomeFila] = fila{componenti: l, indiceCacofonia: 0}
	}
}
