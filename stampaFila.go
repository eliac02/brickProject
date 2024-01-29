package main

import "fmt"

func stampaFila(g gioco, sigma string) {
	if !controllaScatola(g.scatola, sigma) && controllaTavolo(g.tavolo, sigma) {
		_, row := trovaFilaSulTavolo(g.tavolo, sigma)
		fmt.Println("(")
		printList(row.componenti)
		fmt.Println(")")
	}
}
