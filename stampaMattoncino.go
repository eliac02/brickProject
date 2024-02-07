package main

import (
	"fmt"
)

// stampaMattoncino prints the brick named sigma.
//
// @param g The game object.
// @param sigma The name of the brick.
func stampaMattoncino(g gioco, sigma string) {
	if controllaScatola(g.scatola, sigma) {
		fmt.Printf("%s: %s, %s", sigma, g.scatola[sigma][0], g.scatola[sigma][1])
	} else if controllaTavolo(g.tavolo, sigma) {
		_, row := trovaFilaSulTavolo(g.tavolo, sigma)
		nodo := searchNode(row.componenti, sigma)

		fmt.Printf("%s: %s, %s", sigma, nodo.data.alpha, nodo.data.beta)
	}
}
