package main

// inserisciMattoncino inserts a brick into the game.
// It uses linear search because the map g.scatola it's not sorted.
//
// @param g The game object.
// @param alpha The left shape of the brick.
// @param beta The right shape of the brick.
// @param sigma The name of the brick.
//
// If sigma already exists in the box or if alpha equals beta, the function does nothing.
// Otherwise, it inserts the brick alpha beta sigma into the box.
func inserisciMattoncino(g gioco, alpha, beta, sigma string) {
	if alpha != beta {
		for key := range g.scatola {
			if key == sigma {
				return
			}
		}
		g.scatola[sigma] = [2]string{alpha, beta}
		g.forme[alpha] = true
		g.forme[beta] = true
	}
}
