package main

// eliminaFila removes a row of bricks from the game.
//
// If the row exists on the table, its bricks are placed back into the box,
// and the row is removed from the table.
//
// @param g The game object.
// @param sigma The name of the row to be removed.
func eliminaFila(g gioco, sigma string) {
	if !controllaScatola(g.scatola, sigma) && controllaTavolo(g.tavolo, sigma) {
		rowName, row := trovaFilaSulTavolo(g.tavolo, sigma)

		current := row.componenti.head
		for current != nil {
			g.scatola[current.data.sigma] = [2]string{current.data.alpha, current.data.beta}
			g.forme[current.data.alpha] = true
			g.forme[current.data.beta] = true
			current = current.next
		}
		delete(g.tavolo, rowName)
	}
}
