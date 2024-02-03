package main

func indiceCacofonia(g gioco, sigma string) int {
	var result int
	if !controllaScatola(g.scatola, sigma) && controllaTavolo(g.tavolo, sigma) {
		_, row := trovaFilaSulTavolo(g.tavolo, sigma)
		current := row.componenti.head
		for current != nil && current.next != nil {
			lcs := sottostringaMassima(current.data.sigma, current.next.data.sigma)
			row.indiceCacofonia += len(lcs)
			current = current.next
		}
		result = row.indiceCacofonia
	}
	return result
}
