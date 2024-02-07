package main

import "fmt"

// stampaFila prints a row of bricks in the game.
//
// @param g The game object.
// @param sigma The name of the row.
func stampaFila(g gioco, sigma string) {
	if !controllaScatola(g.scatola, sigma) && controllaTavolo(g.tavolo, sigma) {
		_, row := trovaFilaSulTavolo(g.tavolo, sigma)
		fmt.Println("(")
		printList(row.componenti)
		fmt.Println(")")
	}
}
