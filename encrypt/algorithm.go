package encrypt

func Nimbus(text string) string {
	encryptedStr := ""
	for _, alphabet := range text {
		asciicode := int(alphabet)
		newAlphabet := string(asciicode + 3)
		encryptedStr += newAlphabet
	}
	return encryptedStr
}
