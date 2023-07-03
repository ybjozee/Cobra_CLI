package bacon

import (
	"strings"
	"unicode"
)

var lookup = map[string]string{
	"A": "aaaaa", "B": "aaaab", "C": "aaaba", "D": "aaabb", "E": "aabaa",
	"F": "aabab", "G": "aabba", "H": "aabbb", "I": "abaaa", "J": "abaab",
	"K": "ababa", "L": "ababb", "M": "abbaa", "N": "abbab", "O": "abbba",
	"P": "abbbb", "Q": "baaaa", "R": "baaab", "S": "baaba", "T": "baabb",
	"U": "babaa", "V": "babab", "W": "babba", "X": "babbb", "Y": "bbaaa",
	"Z": "bbaab",
}

var reverseLookup = map[string]string{
	"aaaaa": "A", "aaaab": "B", "aaaba": "C", "aaabb": "D", "aabaa": "E", "aabab": "F", "aabba": "G", "aabbb": "H",
	"abaaa": "I", "abaab": "J", "ababa": "K", "ababb": "L", "abbaa": "M", "abbab": "N", "abbba": "O", "abbbb": "P",
	"baaaa": "Q", "baaab": "R", "baaba": "S", "baabb": "T", "babaa": "U", "babab": "V", "babba": "W", "babbb": "X",
	"bbaaa": "Y", "bbaab": "Z",
}

const chunkSize = 5

func Encrypt(plaintext string) string {
	var ciphertext string
	for _, character := range plaintext {
		if cipher, ok := lookup[strings.ToUpper(string(character))]; ok {
			if unicode.IsUpper(character) {
				ciphertext += strings.ToUpper(cipher)
			} else {
				ciphertext += cipher
			}
		} else {
			ciphertext += string(character)
		}
	}
	return ciphertext
}

func Decrypt(ciphertext string) string {
	var plaintext string
	var chunk string
	for i := 0; i < len(ciphertext)/chunkSize; i++ {
		for j := 0; j < chunkSize; j++ {
			chunk += string(ciphertext[(i*chunkSize)+j])
		}
		if plainCharacter, ok := reverseLookup[strings.ToLower(chunk)]; ok {
			if strings.ToLower(chunk) == chunk {
				plaintext += strings.ToLower(plainCharacter)
			} else {
				plaintext += plainCharacter
			}
		} else {
			plaintext += chunk
		}
		chunk = ""
	}
	return plaintext
}
