package main

import (
	"strings"
)

// controllaScatola checks if the brick named sigma is in the box.
//
// @param scatola The box containing the bricks.
// @param sigma The name of the brick.
// @return true if the brick is in the box, false otherwise.
func controllaScatola(scatola map[string][2]string, sigma string) bool {
	if _, exists := scatola[sigma]; exists {
		return true
	} else {
		return false
	}
}

// controllaTavolo checks if the brick named sigma is on the table.
//
// @param tavolo The table containing rows of bricks.
// @param sigma The name of the brick.
// @return true if the brick is on the table, false otherwise.
func controllaTavolo(tavolo map[string]fila, sigma string) bool {
	for key := range tavolo {
		if strings.Contains(key, sigma) {
			return true
		}
	}
	return false
}

// trovaFilaSulTavolo finds the row on the table that contains the brick named sigma.
//
// @param tavolo The table containing rows of bricks.
// @param sigma The name of the brick.
// @return The name of the row containing the brick and the row itself.
func trovaFilaSulTavolo(tavolo map[string]fila, sigma string) (string, fila) {
	for key, value := range tavolo {
		if strings.Contains(key, sigma) {
			return key, value
		}
	}
	return "", fila{}
}

// verificaFila checks if the bricks in the given slice form a valid row.
//
// @param sliceNames The names of the bricks that will form the row.
// @param g The game object.
// @return true if the bricks form a valid row, false otherwise.
func verificaFila(sliceNames []string, g gioco) bool {
	signPrev, sigmaPrev := sliceNames[0][0], sliceNames[0][1:]
	for i := 1; i < len(sliceNames)-1; i++ {
		sign, sigma := sliceNames[i][0], sliceNames[i][1:]
		switch {
		case sign == '+' && signPrev == '+':
			if g.scatola[sigma][0] != g.scatola[sigmaPrev][1] {
				return false
			}
		case sign == '+' && signPrev == '-':
			if g.scatola[sigma][0] != g.scatola[sigmaPrev][0] {
				return false
			}
		case sign == '-' && signPrev == '+':
			if g.scatola[sigma][1] != g.scatola[sigmaPrev][1] {
				return false
			}
		case sign == '-' && signPrev == '-':
			if g.scatola[sigma][1] != g.scatola[sigmaPrev][0] {
				return false
			}
		}
		signPrev, sigmaPrev = sign, sigma
	}
	return true
}
