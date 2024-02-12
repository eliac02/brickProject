//CORTESI ELIA 01911A

package main

import (
	"fmt"
	"math"
)

// contains checks if a given string is present in an array of strings.
//
// The complexity of this function is O(n) where n is the length of the array.
//
// @param array The array of strings to be checked.
// @param str The string to be searched for.
// @return true if the string is present in the array, false otherwise.
func contains(array []string, str string) bool {
	for _, element := range array {
		if element == str {
			return true
		}
	}
	return false
}

// popolaListaAdiacenza populates the adjacency list with adjacent bricks for each shape in the game.
//
// The complexity of this function is O(n*m*k) where n is the number of shapes, m is the number of bricks in the game and k is the length of the adjacency list.
//
// @param g The game state.
// @param listaAdiacenza The adjacency list to be populated.
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

// creaListaNomi creates a string containing the names of bricks forming a row.
//
// The complexity of this function is O(n*m) where n is the length of sequenzaForme and m is the number of bricks in the game.
//
// @param g The game state.
// @param sequenzaForme The sequence of shapes.
// @return A string containing the names of the bricks forming the row.
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

// ricostruisciSequenza reconstructs the sequence of shapes from beta to alpha using predecessor mapping.
//
// The complexity of this function is O(n) where n is the length of the predecessor mapping.
//
// @param alpha The starting shape.
// @param beta The ending shape.
// @param predecessori The predecessor mapping.
// @return The sequence of shapes from beta to alpha.
func ricostruisciSequenza(alpha, beta string, predecessori map[string]string) []string {
	var sequenzaForme []string
	for current := beta; current != alpha; current = predecessori[current] {
		sequenzaForme = append(sequenzaForme, current)
	}
	sequenzaForme = append(sequenzaForme, alpha)
	return sequenzaForme
}

// controllaCasoSemplice checks if a minimum path of lenght three exists between to adjacent shapes.
//
// The complexity of this function is O(n*m) where n is the number of shapes and m is the number of bricks in the game.
//
// @param g The game state.
// @param listaAdiacenza The adjacency list.
// @param alpha The starting shape.
// @return true if a minimum path of lenght three exists between to adjacent shapes, along with the sequence of bricks forming the path, otherwise false and an empty string.
func controllaCasoSemplice(g gioco, listaAdiacenza map[string][]string, alpha string) (bool, string) {
	listaNomi := ""
	counter := 0
	//controllo tra i vicini di alpha
	for _, adj := range listaAdiacenza[alpha] {
		//controlla se ci sono i mattoncini giusti nella scatola
		for nome, bordi := range g.scatola {
			if bordi[0] == alpha && bordi[1] == adj {
				if counter == 1 {
					listaNomi += "-" + nome + " "
				} else {
					listaNomi += "+" + nome + " "
				}
				counter++
			} else if bordi[1] == alpha && bordi[0] == adj {
				if counter == 1 {
					listaNomi += "+" + nome + " "
				} else {
					listaNomi += "-" + nome + " "
				}
				counter++
			}
			if counter == 2 {
				return true, listaNomi
			}
		}
		listaNomi = ""
		counter = 0
	}
	return false, ""
}

// eliminaCollegamento removes the connection between two shapes in the adjacency list.
//
// The complexity of this function is O(n) where n is the length of the adjacency list.
//
// @param listaAdiacenza The adjacency list.
// @param alpha The first shape.
// @param beta The second shape.
func eliminaCollegamento(listaAdiacenza map[string][]string, alpha, beta string) {
	//elimino i collegamenti tra alpha e beta
	for i, adj := range listaAdiacenza[alpha] {
		if adj == beta {
			listaAdiacenza[alpha] = append(listaAdiacenza[alpha][:i], listaAdiacenza[alpha][i+1:]...)
		}
	}
	for i, adj := range listaAdiacenza[beta] {
		if adj == alpha {
			listaAdiacenza[beta] = append(listaAdiacenza[beta][:i], listaAdiacenza[beta][i+1:]...)
		}
	}
}

// aggiungiCollegamento adds a connection between two shapes in the adjacency list.
//
// The complexity of this function is O(n) where n is the length of the adjacency list.
//
// @param listaAdiacenza The adjacency list.
// @param alpha The first shape.
// @param beta The second shape.
func aggiungiCollegamento(listaAdiacenza map[string][]string, alpha, beta string) {
	//riaggiungo i collegamenti tra alpha e beta
	listaAdiacenza[alpha] = append(listaAdiacenza[alpha], beta)
	listaAdiacenza[beta] = append(listaAdiacenza[beta], alpha)
}

// bfsAlphaUgualebeta performs a Breadth-First Search (BFS) to find the shortest path between two equal shapes.
//
// The compexity of this function is O(n+m), where n is the number of shapes and m is the number of bricks in the game.
//
// @param listaAdiacenza The adjacency list.
// @param alpha The starting shape.
// @param beta The ending shape.
// @return The shortest sequence of shapes from alpha to beta.
func bfsAlphaUgualebeta(listaAdiacenza map[string][]string, alpha, beta string) []string {
	//creo tutte le strutture dati necessarie

	var camminoMinimo []string
	var lenCamminoMinimo int = math.MaxInt32

	for _, adj := range listaAdiacenza[alpha] {
		//elimino i collegamenti tra adj e alpha
		eliminaCollegamento(listaAdiacenza, alpha, adj) //O(n)
		//trovo il cammino tra adj e beta
		sequenzaForme := bfsNormale(listaAdiacenza, adj, beta)
		if lenCamminoMinimo > len(sequenzaForme) {
			lenCamminoMinimo = len(sequenzaForme)
			camminoMinimo = sequenzaForme
		}
		//aggiungo il collegamento tra adj e beta
		aggiungiCollegamento(listaAdiacenza, alpha, adj) //O(n)
	}

	return append(camminoMinimo, alpha)
}

// bfsNormale performs a Breadth-First Search (BFS) to find the shortest path between two different shapes.
//
// The compexity of this function is O(n+m), where n is the number of shapes and m is the number of bricks in the game.
//
// @param listaAdiacenza The adjacency list.
// @param alpha The starting shape.
// @param beta The ending shape.
// @return The shortest sequence of shapes from alpha to beta.
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

// disponiFilaMinima arranges the bricks in the game to form the shortest sequence from alpha to beta.
//
// The complexity of this function is O(n*(m^2)) where m is the number of shapes and n is the number of bricks in the game.
//
// @param g The game object.
// @param alpha The starting shape.
// @param beta The ending shape.
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
			sequenzaForme := bfsAlphaUgualebeta(listaAdiacenza, alpha, beta)
			listaNomi := creaListaNomi(g, sequenzaForme)

			//controllo se listaNomi e' vuota
			if listaNomi == "" {
				fmt.Printf("non esiste fila da %s a %s\n", alpha, beta)
			} else {
				//elimino l'ultimo spazio e dispongo la fila
				listaNomi = listaNomi[:len(listaNomi)-1]
				disponiFila(g, listaNomi)
			}
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
