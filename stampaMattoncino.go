//CORTESI ELIA 01911A

package main

import (
	"fmt"
)

// stampaMattoncino prints the brick named sigma.
//
// The complexity of this function is O(n*m*k), where n is the length of the row's name, m is the length of the brick's name and k is the number of rows.
//
// @param g The game object.
// @param sigma The name of the brick.
func stampaMattoncino(g gioco, sigma string) {
	if controllaScatola(g.scatola, sigma) {
		fmt.Printf("%s: %s, %s", sigma, g.scatola[sigma][0], g.scatola[sigma][1])
	} else if controllaTavolo(g.tavolo, sigma) {
		_, row := trovaFilaSulTavolo(g.tavolo, sigma)
		_, nodo := searchNode(row.componenti, sigma)

		fmt.Printf("%s: %s, %s", sigma, nodo.data.alpha, nodo.data.beta)
	}
}
