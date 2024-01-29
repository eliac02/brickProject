package main

func eliminaFila(g gioco, sigma string) {
	if !controllaScatola(g.scatola, sigma) && controllaTavolo(g.tavolo, sigma) {
		rowName, row := trovaFilaSulTavolo(g.tavolo, sigma)

		current := row.componenti.head
		for current != nil {
			g.scatola[current.data.sigma] = [2]string{current.data.alpha, current.data.beta}
			current = current.next
		}
		delete(g.tavolo, rowName)
	}
}
