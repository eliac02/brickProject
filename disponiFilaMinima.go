package main

import (
	"fmt"
)

type arco struct {
	start string
	end   string
}

func disponiFilaMinima(g gioco, alpha, beta string) {
	queue := newQueue()
	elencoArchi := make(map[arco]string)
	elencoPredecessori := make(map[string]string)
	visited := make(map[string]bool)

	//riempio elencoArchi con le forme dei mattoncini presenti nella scatola
	for nome, bordi := range g.scatola {
		elencoArchi[arco{bordi[0], bordi[1]}] = nome
	}

	if alpha == beta {
		//implementare la funzione nel caso in cui alpha e beta siano uguali
	} else {
		enqueue(queue, alpha)
		elencoPredecessori[alpha] = ""

		//esploro il grafo
		for !isEmpty(queue) {
			current := bottom(queue)
			dequeue(queue)
			visited[current] = true
			for arco := range elencoArchi {
				if arco.start == current && !visited[arco.end] {
					enqueue(queue, arco.end)
					visited[arco.end] = true
					elencoPredecessori[arco.end] = arco.start
				} else if arco.end == current && !visited[arco.start] {
					enqueue(queue, arco.start)
					visited[arco.start] = true
					elencoPredecessori[arco.start] = arco.end
				}
			}
		}

		//ricostruisco la sequenza di forme
		var sequenzaForme []string
		for current := beta; current != ""; current = elencoPredecessori[current] {
			sequenzaForme = append(sequenzaForme, current)
		}

		// creare elenco dei mattoncini
		var listaNomi string
		for i := len(sequenzaForme) - 1; i > 0; i-- {
			// devo controllare entrambi i campi di arco all'interno di elencoArchi
			arcoCorrente := arco{sequenzaForme[i], sequenzaForme[i-1]}
			if nome, exists := elencoArchi[arcoCorrente]; exists {
				brick := "+" + nome + " "
				listaNomi += brick
			} else {
				arcoCorrente = arco{sequenzaForme[i-1], sequenzaForme[i]}
				nome := elencoArchi[arcoCorrente]
				brick := "-" + nome + " "
				listaNomi += brick
			}
		}
		//stampo elencoArchi
		if listaNomi == "" {
			fmt.Printf("non esiste fila da %s a %s\n", alpha, beta)
			return
		}
		listaNomi = listaNomi[:len(listaNomi)-1]

		//dispongo fila
		disponiFila(g, listaNomi)
	}

}
