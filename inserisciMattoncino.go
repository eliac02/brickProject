package main

// uso una ricerca lineare perche' ho frequenti inserimenti
func inserisciMattoncino(g gioco, alpha, beta, sigma string) {
	//se esiste gia' sigma nella scatola, oppure se alpha = beta, non faccio nulla
	//altrimenti inserisco nella scatola il mattoncino alpha beta sigma
	if alpha != beta {
		if len(g.scatola) > 0 {
			for key := range g.scatola {
				if key == sigma {
					return
				}
			}
		}
		g.scatola[sigma] = [2]string{alpha, beta}
		g.forme[alpha] = true
		g.forme[beta] = true
	}
}
