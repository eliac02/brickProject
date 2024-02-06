package main

import "fmt"

func contains(array []string, str string) bool {
	for _, element := range array {
		if element == str {
			return true
		}
	}
	return false
}

func popolaListaAdiacenza(g gioco, listaAdiacenza map[string][]string) {
	for forma := range g.forme {
		listaAdiacenza[forma] = []string{}
		for _, mattoncino := range g.scatola {
			if mattoncino[0] == forma && !contains(listaAdiacenza[forma], mattoncino[1]) {
				listaAdiacenza[forma] = append(listaAdiacenza[forma], mattoncino[1])
			}
			if mattoncino[1] == forma && !contains(listaAdiacenza[forma], mattoncino[0]) {
				listaAdiacenza[forma] = append(listaAdiacenza[forma], mattoncino[0])
			}
		}
	}
}

func creaListaNomi(g gioco, sequenzaForme []string) string {
	var listaNomi string
	for i := len(sequenzaForme) - 1; i > 0; i-- {
		//controllo i bordi dei mattoncini nella scatola
		//se combaciano ne concateno il nome alla stringa
		for nome, bordi := range g.scatola {
			if sequenzaForme[i] == bordi[0] && sequenzaForme[i-1] == bordi[1] {
				listaNomi += "+" + nome + " "
			}
			if sequenzaForme[i] == bordi[1] && sequenzaForme[i-1] == bordi[0] {
				listaNomi += "-" + nome + " "
			}
		}
	}
	return listaNomi
}

func ricostruisciSequenza(alpha, beta string, predecessori map[string]string) []string {
	var sequenzaForme []string
	for current := beta; current != alpha; current = predecessori[current] {
		sequenzaForme = append(sequenzaForme, current)
	}
	sequenzaForme = append(sequenzaForme, alpha)
	return sequenzaForme
}

func controllaCasoSemplice(g gioco, listaAdiacenza map[string][]string, alpha string) (bool, string) {
	listaNomi := ""
	counter := 0
	//controllo tra i vicini di alpha
	for _, adj := range listaAdiacenza[alpha] {
		//controlla se ci sono i mattoncini giusti nella scatola
		for nome, bordi := range g.scatola {
			if bordi[0] == alpha && bordi[1] == adj {
				listaNomi += "+" + nome + " "
				counter++
			} else if bordi[1] == alpha && bordi[0] == adj {
				listaNomi += "-" + nome + " "
				counter++
			}
		}
		if counter == 2 {
			return true, listaNomi
		} else {
			listaNomi = ""
			counter = 0
		}
	}
	return false, ""
}

func bfsNormale(listaAdiacenza map[string][]string, alpha, beta string) []string {
	//creo tutte le strutture dati necessarie

	// Mappa per tenere traccia dei nodi visitati
	nodiVisitati := make(map[string]bool)
	//Mappa dei predecessori
	predecessori := make(map[string]string)
	// Coda per la BFS
	coda := newQueue()

	// Aggiungo il nodo di partenza alla coda e lo marco come visitato
	coda.enqueue(alpha)
	nodiVisitati[alpha] = true

	// Finché ci sono nodi nella coda
	for !coda.isEmpty() {
		currentNode := coda.bottom()
		coda.dequeue()

		//visito i nodi adiacenti
		for _, adj := range listaAdiacenza[currentNode] {
			if !nodiVisitati[adj] {
				nodiVisitati[adj] = true
				coda.enqueue(adj)
				predecessori[adj] = currentNode

				//se trovo il nodo beta esco
				if adj == beta {
					return ricostruisciSequenza(alpha, beta, predecessori)
				}
			}
		}
	}
	return nil
}

func disponiFilaMinima(g gioco, alpha, beta string) {
	//Creo la mappa di liste di adiacenza e la popolo
	listaAdiacenza := make(map[string][]string)
	popolaListaAdiacenza(g, listaAdiacenza)

	if alpha == beta {
		//controllo se e' il caso facile o difficile
		easy, listaNomi := controllaCasoSemplice(g, listaAdiacenza, alpha)
		if easy {
			//elimino l'ultimo spazio e dispongo la fila
			listaNomi = listaNomi[:len(listaNomi)-1]
			disponiFila(g, listaNomi)
		} else {
			//fare caso in cui alpha != beta
			fmt.Println("hello world")
		}

	} else {
		sequenzaForme := bfsNormale(listaAdiacenza, alpha, beta)
		listaNomi := creaListaNomi(g, sequenzaForme)

		//controllo se listaNomi e' vuota
		if listaNomi == "" {
			fmt.Printf("non esiste fila da %s a %s\n", alpha, beta)
		} else {
			//elimino l'ultimo spazio e disponog la fila
			listaNomi = listaNomi[:len(listaNomi)-1]
			disponiFila(g, listaNomi)
		}
	}
}
