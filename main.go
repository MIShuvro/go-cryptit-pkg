package main

import (
	"fmt"
	"github.com/MIShuvro/go-cryptit-pkg/decrypt"
	"github.com/MIShuvro/go-cryptit-pkg/encrypt"
)

func main() {

	encryptedText := encrypt.Nimbus("text")

	fmt.Println("encrypted text==========>", encryptedText)
	decryptedText := decrypt.Nimbus(encryptedText)
	fmt.Println("decrypted text==========>", decryptedText)
}
