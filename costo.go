package main

import (
	"fmt"
	"strings"
)

func min3(a, b, c int) int {
	minimum := a
	if b < minimum {
		minimum = b
	}
	if c < minimum {
		minimum = c
	}
	return minimum
}

func min2(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// crea l'elenco dei nomi (o nella scatola o nella fila) dalle forme
func trovaElencoMattoncini(g gioco, row fila, sequenzaForme []string) []string {
	elencoNomi := []string{}
	for i := 0; i < len(sequenzaForme)-1; i++ {
		//devo controllare se la coppia di forme è presente nella scatola oppure se e' un mattoncino già presente in row
		for nome, bordi := range g.scatola {
			if sequenzaForme[i] == bordi[0] && sequenzaForme[i+1] == bordi[1] {
				elencoNomi = append(elencoNomi, nome)
			} else if sequenzaForme[i] == bordi[1] && sequenzaForme[i+1] == bordi[0] {
				elencoNomi = append(elencoNomi, nome)
			} else {
				//scorro la fila per vedere se la coppia di forme corrisponde ad un mattoncino già presente
				current := row.componenti.head
				for current != nil {
					if sequenzaForme[i] == current.data.alpha && sequenzaForme[i+1] == current.data.beta {
						elencoNomi = append(elencoNomi, current.data.sigma)
					} else if sequenzaForme[i] == current.data.beta && sequenzaForme[i+1] == current.data.alpha {
						elencoNomi = append(elencoNomi, current.data.sigma)
					}
					current = current.next
				}
			}
		}
	}
	return elencoNomi
}

func trovaFormeIniziali(g gioco, row fila) []string {
	elencoForme := []string{}
	current := row.componenti.head
	for current != nil {
		if len(elencoForme) == 0 {
			elencoForme = append(elencoForme, current.data.alpha)
			elencoForme = append(elencoForme, current.data.beta)
		} else {
			elencoForme = append(elencoForme, current.data.beta)
		}
		current = current.next
	}
	return elencoForme
}

func verificaMattonciniDisponibili(g gioco, row fila, elencoNomiFinale []string) bool {
	for _, name := range elencoNomiFinale {
		isPresent, _ := searchNode(row.componenti, name)
		if !controllaScatola(g.scatola, name) && !isPresent {
			fmt.Println("indefinito")
			return false
		}
	}
	return true
}

func costo(g gioco, sigma string, nomiFinali string) {
	_, row := trovaFilaSulTavolo(g.tavolo, sigma)
	sequenzaNomiFinali := trovaElencoMattoncini(g, row, strings.Split(nomiFinali, " "))
	if (len(sequenzaNomiFinali) != len(nomiFinali)-1) || !verificaMattonciniDisponibili(g, row, sequenzaNomiFinali) {
		fmt.Println("indefinito")
	}
	listaFormeIniziali := trovaFormeIniziali(g, row)
	listaFormeFinali := strings.Split(nomiFinali, " ")

	matrix := make([][]int, len(listaFormeIniziali)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(listaFormeFinali)+1)
	}

	for i := 0; i <= len(listaFormeIniziali); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(listaFormeFinali); j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len(listaFormeIniziali); i++ {
		for j := 1; j <= len(listaFormeFinali); j++ {
			if listaFormeIniziali[i-1] == listaFormeFinali[j-1] {
				opCost := min3(matrix[i-1][j]+1, matrix[i][j-1]+1, matrix[i-1][j-1])
				matrix[i][j] = opCost
			} else {
				opCost := min2(matrix[i-1][j]+1, matrix[i][j-1]+1)
				matrix[i][j] = opCost
			}

		}
	}
	fmt.Println(matrix[len(listaFormeIniziali)][len(listaFormeFinali)])
}
