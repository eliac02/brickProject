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
		//aggiunge forma alla lista di adiacenza
		listaAdiacenza[forma] = make([]string, 0)
		//controlla se la forma e' adiacente ad un'altra forma
		//per farlo controllo le forme dei mattoncini dentro alla scatola
		for _, mattoncino := range g.scatola {
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
		sequenza = append(sequenza, current)
	}
	return sequenza
}

func disponiFilaMinima(g gioco, alpha, beta string) {
	//modello il grafo come una lista di adiacenza
	listaAdiacenza := make(map[string][]string)
	coda := newQueue()
	elencoNodiVisitati := make(map[string]bool)
	elencoPredecessori := make(map[string]string)

	popolaListaAdiacenza(g, listaAdiacenza)

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
		//controllare cosa succede se non c'e' la sequenza
	}

	sequenzaForme := ricostruisciSequenza(alpha, beta, elencoPredecessori)

	//ottieni una stringa con le forme da dare a disponiFila()
	//ciclo su sequenzaForme e aggiungo le forme alla stringa con il segno corretto
	listaNomi := ""
	for i := len(sequenzaForme) - 1; i > 0; i-- {
		//guardo le forme a coppie e controllo dentro a g.scatola quale sia il nome del mattoncino
		//se le due forme sono i bordi di un mattoncino concateno il nome del mattoncino alla stringa
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

	//controllo che listaNomi non sia vuota
	if listaNomi == "" {
		fmt.Printf("non esiste fila da %s a %s\n", alpha, beta)
		return
	}
	//toilgo lo spazio finale
	listaNomi = listaNomi[:len(listaNomi)-1]

	//dispongo la fila
	disponiFila(g, listaNomi)
}
