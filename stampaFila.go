//CORTESI ELIA 01911A

package main

import "fmt"

// stampaFila prints a row of bricks in the game.
//
// The complexity of this function is O(n*m*k), where n is the length of the row's name, m is the length of the brick's name and k is the number of rows.
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
