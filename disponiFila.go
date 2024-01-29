package main

import "strings"

func creaFila(g gioco, tempSlice []string, l *linkedList) {
	for _, name := range tempSlice {
		node := newNode(mattoncino{alpha: g.scatola[name][0], beta: g.scatola[name][1], sigma: name})
		addNode(l, node)
	}
}

func disponiFila(g gioco, listaNomi string) {
	var nomeFila string
	sliceNames := strings.Split(listaNomi, " ")

	//controlla che tutti i mattoncini siano nella scatola
	for _, name := range sliceNames {
		temp := name[1:]
		if !controllaScatola(g.scatola, temp) {
			return
		}
		nomeFila += temp + " "
	}

	//controlla che i mattoncni formino una fila
	verificaFila(sliceNames, g)

	//dispongo la fila sul tavolo
	tempSlice := strings.Split(nomeFila, " ")
	l := newList()
	creaFila(g, tempSlice, l)
	g.tavolo[nomeFila] = l
}
