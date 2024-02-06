package main

import (
	"fmt"
)

func contains(array []string, str string) bool {
	for _, element := range array {
		if element == str {
			return true
		}
	}
	return false
}

func rimuoviElementoDaArray(array []string, indice int) []string {
	return append(array[:indice], array[indice+1:]...)
}

func popolaListaAdiacenza(g gioco, listaAdiacenza map[string][]string) { //O(n^3)
	for forma := range g.forme { //O(m)
		//aggiunge forma alla lista di adiacenza
		listaAdiacenza[forma] = make([]string, 0)
		//controlla se la forma e' adiacente ad un'altra forma
		//per farlo controllo le forme dei mattoncini dentro alla scatola
		for _, mattoncino := range g.scatola { //O(n)
			//controllo anche che la forma non sia gia' presente nella lista di adiacenza
			if mattoncino[0] == forma && !contains(listaAdiacenza[forma], mattoncino[1]) {
				listaAdiacenza[forma] = append(listaAdiacenza[forma], mattoncino[1])
			} else if mattoncino[1] == forma && !contains(listaAdiacenza[forma], mattoncino[0]) {
				listaAdiacenza[forma] = append(listaAdiacenza[forma], mattoncino[0])
			}
		}
	}
}

func ricostruisciSequenza(alpha, beta string, elencoPredecessori map[string]string) []string {
	var sequenza []string
	for current := beta; current != ""; current = elencoPredecessori[current] {
		sequenza = append([]string{current}, sequenza...)
	}
	return sequenza
}

func creaListaNomi(g gioco, sequenzaForme []string) string {
	listaNomi := ""
	for i := len(sequenzaForme) - 1; i > 0; i-- {
		for nome, bordi := range g.scatola {
			coppiaForme := [2]string{sequenzaForme[i], sequenzaForme[i-1]}
			if bordi == coppiaForme {
				listaNomi += "+" + nome + " "
				break
			}
			coppiaFormeInversa := [2]string{sequenzaForme[i-1], sequenzaForme[i]}
			if bordi == coppiaFormeInversa {
				listaNomi += "-" + nome + " "
				break
			}
		}
	}
	return listaNomi
}

func trovaSequenzaPiuCorta(elencoSequenze [][]string) int {
	min := len(elencoSequenze[0])
	var i int
	for index, sequenza := range elencoSequenze {
		if len(sequenza) < min {
			min = len(sequenza)
			i = index
		}
	}
	return i
}

func casoAlphaUgualeBeta(g gioco, coda *queue, listaAdiacenza map[string][]string, elencoPredecessori map[string]string, elencoNodiVisitati map[string]bool, alpha, beta string) []string {
	var elencoSequenze [][]string
	var sequenzaForme []string
	for _, adj := range listaAdiacenza[alpha] {
		tempMap := make(map[string][]string)
		for k, v := range listaAdiacenza {
			tempMap[k] = append([]string{}, v...)
		}
		for i, v := range tempMap[adj] {
			if v == alpha {
				tempArray := rimuoviElementoDaArray(tempMap[adj], i)
				tempMap[adj] = tempArray
			}
		}
		for i, v := range tempMap[alpha] {
			if v == adj {
				tempArray := rimuoviElementoDaArray(tempMap[alpha], i)
				tempMap[alpha] = tempArray
			}
		}
		sequenzaForme = bfsNormale(g, coda, tempMap, elencoPredecessori, elencoNodiVisitati, adj, beta)
		elencoSequenze = append(elencoSequenze, sequenzaForme)
	}
	index := trovaSequenzaPiuCorta(elencoSequenze)
	return elencoSequenze[index]
}

func bfsNormale(g gioco, coda *queue, listaAdiacenza map[string][]string, elencoPredecessori map[string]string, elencoNodiVisitati map[string]bool, alpha, beta string) []string {
	coda.enqueue(alpha)
	elencoPredecessori[alpha] = ""
	elencoNodiVisitati[alpha] = true

	for !coda.isEmpty() {
		current := coda.bottom()
		coda.dequeue()
		for _, forma := range listaAdiacenza[current] {
			if !elencoNodiVisitati[forma] {
				coda.enqueue(forma)
				elencoPredecessori[forma] = current
				elencoNodiVisitati[forma] = true
			}
		}
	}

	return ricostruisciSequenza(alpha, beta, elencoPredecessori)
}

func disponiFilaMinima(g gioco, alpha, beta string) {
	listaAdiacenza := make(map[string][]string)
	coda := newQueue()
	elencoNodiVisitati := make(map[string]bool)
	elencoPredecessori := make(map[string]string)

	popolaListaAdiacenza(g, listaAdiacenza)

	if alpha == beta {
		sequenzaForme := append([]string{}, alpha)
		tempSequenzaForme := casoAlphaUgualeBeta(g, coda, listaAdiacenza, elencoPredecessori, elencoNodiVisitati, alpha, beta)
		sequenzaForme = append(sequenzaForme, tempSequenzaForme...)
		listaNomi := creaListaNomi(g, sequenzaForme)
		if listaNomi == "" {
			fmt.Printf("non esiste fila da %s a %s\n", alpha, beta)
			return
		}
		listaNomi = listaNomi[:len(listaNomi)-1]
		disponiFila(g, listaNomi)
	} else {
		sequenzaForme := bfsNormale(g, coda, listaAdiacenza, elencoPredecessori, elencoNodiVisitati, alpha, beta)
		listaNomi := creaListaNomi(g, sequenzaForme)
		if listaNomi == "" {
			fmt.Printf("non esiste fila da %s a %s\n", alpha, beta)
			return
		}
		listaNomi = listaNomi[:len(listaNomi)-1]
		disponiFila(g, listaNomi)
	}
}
