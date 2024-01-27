package main

import (
	"fmt"
	"strings"
)

// controlla se il mattoncino sigma e' nella scatola
func controllaScatola(scatola map[string][2]string, sigma string) bool {
	if _, exists := scatola[sigma]; exists {
		return true
	} else {
		return false
	}
}

// controlla se il mattoncino sigma e' sul tavolo
func controllaTavolo(tavolo map[string]*linkedList, sigma string) bool {
	for key := range tavolo {
		if strings.Contains(key, sigma) {
			return true
		}
	}
	return false
}

// trova la fila sul tavolo che contiene il mattoncino sigma
func trovaFilaSulTavolo(tavolo map[string]*linkedList, sigma string) *linkedList {
	for key := range tavolo {
		if strings.Contains(key, sigma) {
			return tavolo[key]
		}
	}
	return nil
}

// stampa il mattoncino sigma
func stampaMattoncino(g gioco, sigma string) {
	if controllaScatola(g.scatola, sigma) {
		fmt.Printf("%s: %s, %s", sigma, g.scatola[sigma][0], g.scatola[sigma][1])
	} else if controllaTavolo(g.tavolo, sigma) {
		row := trovaFilaSulTavolo(g.tavolo, sigma)
		nodo := searchNode(row, sigma)
		fmt.Printf("%s: %s, %s", sigma, nodo.data.alpha, nodo.data.beta)
	}
}
