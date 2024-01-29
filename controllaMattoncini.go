package main

import (
	"strings"
)

// controlla se il mattoncino sigma e' nella scatola
func controllaScatola(scatola map[string][2]string, sigma string) bool {
	if _, exists := scatola[sigma]; exists {
		return true
	} else {
		return false
	}
}

// controlla se il mattoncino sigma e' sul tavolo
func controllaTavolo(tavolo map[string]fila, sigma string) bool {
	for key := range tavolo {
		if strings.Contains(key, sigma) {
			return true
		}
	}
	return false
}

// trova la fila sul tavolo che contiene il mattoncino sigma
func trovaFilaSulTavolo(tavolo map[string]fila, sigma string) (string, fila) {
	for key, value := range tavolo {
		if strings.Contains(key, sigma) {
			return key, value
		}
	}
	return "", fila{}
}

func verificaFila(sliceNames []string, g gioco) {
	signPrev, sigmaPrev := sliceNames[0][0], sliceNames[0][1:]
	for i := 1; i < len(sliceNames)-1; i++ {
		sign, sigma := sliceNames[i][0], sliceNames[i][1:]
		switch {
		case sign == '+' && signPrev == '+':
			if g.scatola[sigma][0] != g.scatola[sigmaPrev][1] {
				return
			}
		case sign == '+' && signPrev == '-':
			if g.scatola[sigma][0] != g.scatola[sigmaPrev][0] {
				return
			}
		case sign == '-' && signPrev == '+':
			if g.scatola[sigma][1] != g.scatola[sigmaPrev][1] {
				return
			}
		case sign == '-' && signPrev == '-':
			if g.scatola[sigma][1] != g.scatola[sigmaPrev][0] {
				return
			}
		}
		signPrev, sigmaPrev = sign, sigma
	}
}
