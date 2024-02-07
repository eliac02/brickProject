package main

// sottostringaMassima finds the longest common substring between two strings using dynamic programming tecnique.
//
// @param sigma The first string.
// @param teta The second string.
// @return The longest common substring between sigma and teta.
func sottostringaMassima(sigma, teta string) string {
	lcs := make([][]int, len(sigma)+1)
	for i := range lcs {
		lcs[i] = make([]int, len(teta)+1)
	}
	for i := 0; i <= len(sigma); i++ {
		lcs[i][0] = 0
	}
	for j := 0; j <= len(teta); j++ {
		lcs[0][j] = 0
	}
	for i := 1; i <= len(sigma); i++ {
		for j := 1; j <= len(teta); j++ {
			if sigma[i-1] == teta[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	lenMax := lcs[len(sigma)][len(teta)]

	// Se non c'e' nessuna sottostringa comune
	if lenMax == 0 {
		return ""
	}

	// Costruisco la sottostringa massima
	substring := make([]byte, lenMax)
	i, j := len(sigma), len(teta)
	for lenMax > 0 {
		if lcs[i][j] != lcs[i-1][j-1] {
			substring[lenMax-1] = sigma[i-1]
			i--
			j--
			lenMax--
		} else {
			if lcs[i][j] == lcs[i-1][j] {
				i--
			} else {
				j--
			}
		}
	}
	return string(substring)
}
