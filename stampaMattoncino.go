package main

import (
	"fmt"
)

// stampa il mattoncino sigma
func stampaMattoncino(g gioco, sigma string) {
	if controllaScatola(g.scatola, sigma) {
		fmt.Printf("%s: %s, %s", sigma, g.scatola[sigma][0], g.scatola[sigma][1])
	} else if controllaTavolo(g.tavolo, sigma) {
		_, row := trovaFilaSulTavolo(g.tavolo, sigma)
		nodo := searchNode(row.componenti, sigma)
		//stampa effettiva
		fmt.Printf("%s: %s, %s", sigma, nodo.data.alpha, nodo.data.beta)
	}
}
