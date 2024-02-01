package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZÁÉÍÓÚabcdefghijklmnopqrstuvwxyzáéíóú0123456789!\"£$%^&*()[]{};'#:@~<>?`¦¬€-="

func hashedLetter(key int, letter string) (result string) {
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	leftoverLetters := string(runes[0 : len(letter)-key])

	return fmt.Sprintf(`%s%s`, lastLetterKey, leftoverLetters)
}

func encrypt(key int, plainText string) (result string) {
	hashedLetter := hashedLetter(key, originalLetter)

	var hashedString = ""

	findOne := func(r rune) rune {
		pos := strings.Index(originalLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			hashedString = hashedString + string(hashedLetter[letterPosition])
			return r
		}
		return r
	}

	strings.Map(findOne, plainText)
	return hashedString
}

func decrypt(key int, encryptedText string) (result string) {
	hashedLetter := hashedLetter(key, originalLetter)

	var hashedString = ""

	findOne := func(r rune) rune {
		pos := strings.Index(hashedLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			hashedString = hashedString + string(originalLetter[letterPosition])
			return r
		}
		return r
	}

	strings.Map(findOne, encryptedText)
	return hashedString

}

func main() {
	var plainText string

	fmt.Println("Available text:\n", originalLetter)

	fmt.Print("\nPlease type word:\n>")
	_, err := fmt.Scanf("%s", &plainText)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("\nPlain Text: ", plainText)

	encrypted := encrypt(5, plainText)
	fmt.Println("Enypted Text:", encrypted)

	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted: ", decrypted)
}
