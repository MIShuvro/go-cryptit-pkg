package decrypt

func Nimbus(encryptedText string) string {
	decryptText := ""
	for _, alphabet := range encryptedText {
		asciicode := int(alphabet) - 3
		decryptText += string(asciicode)
	}
	return decryptText
}
