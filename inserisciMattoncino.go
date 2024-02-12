//CORTESI ELIA 01911A

package main

// inserisciMattoncino inserts a brick into the game.
//
// The complexity of this function is O(1).
//
// @param g The game object.
// @param alpha The left shape of the brick.
// @param beta The right shape of the brick.
// @param sigma The name of the brick.
func inserisciMattoncino(g gioco, alpha, beta, sigma string) {
	if alpha != beta {
		if controllaScatola(g.scatola, sigma) {
			return
		}
		g.scatola[sigma] = [2]string{alpha, beta}
		g.forme[alpha] = true
		g.forme[beta] = true
	}
}
