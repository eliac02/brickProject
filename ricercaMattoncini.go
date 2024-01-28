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
func controllaTavolo(tavolo map[string]linkedList, sigma string) bool {
	for key := range tavolo {
		if strings.Contains(key, sigma) {
			return true
		}
	}
	return false
}

// trova la fila sul tavolo che contiene il mattoncino sigma
func trovaFilaSulTavolo(tavolo map[string]linkedList, sigma string) (string, linkedList) {
	for key := range tavolo {
		if strings.Contains(key, sigma) {
			return key, tavolo[key]
		}
	}
	return "", nil
}
