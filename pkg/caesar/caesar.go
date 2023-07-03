package caesar

import (
	"strings"
	"unicode"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func Encrypt(plaintext string, rotations int) string {
	var ciphertext string

	for _, character := range plaintext {
		index := getCharacterIndex(strings.ToLower(string(character)))
		if index == -1 {
			ciphertext += string(character)
		} else {
			cipher := alphabets[getCipherIndex(index, rotations)]
			if unicode.IsUpper(character) {
				ciphertext += strings.ToUpper(string(cipher))
			} else {
				ciphertext += string(cipher)
			}
		}
	}
	return ciphertext
}

func Decrypt(ciphertext string, rotations int) string {
	var plaintext string
	for _, character := range ciphertext {
		index := getCharacterIndex(strings.ToLower(string(character)))
		if index == -1 {
			plaintext += string(character)
		} else {
			plainCharacter := alphabets[getPlaintextIndex(index, int(rotations))]
			if unicode.IsUpper(character) {
				plaintext += strings.ToUpper(string(plainCharacter))
			} else {
				plaintext += string(plainCharacter)
			}
		}
	}
	return plaintext
}

func getCharacterIndex(element string) int {
	return strings.Index(alphabets, element)
}

func getCipherIndex(plainIndex int, key int) int {
	newIndex := plainIndex + key
	if newIndex > 25 {
		return newIndex % 26
	}
	return newIndex
}

func getPlaintextIndex(cipherIndex int, key int) int {
	newIndex := cipherIndex - key
	if newIndex < 0 {
		newIndex = 26 + (newIndex % 26)
	}
	if newIndex > 25 {
		return newIndex % 26
	}
	return newIndex
}
