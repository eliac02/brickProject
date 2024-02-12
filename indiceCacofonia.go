//CORESI ELIA 01911A

package main

// indiceCacofonia calculates the cacophony index of a given brick in the game.
//
// The complexity of this function is O(n*m*k), where n is the lenght of row's name, m is the length of the brick's name and k is the number of rows on the table.
//
// @param g The game state.
// @param sigma The name of the brick.
// @return The cacophony index of the row.
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
