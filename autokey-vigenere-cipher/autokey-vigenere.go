/*
	Pending:
	1. untuk fungsi enkripsi dan dekripsi hanya bisa menggunakan string upper case saja, jika lower case dekripsi gagal
	2. perlu membuat fungsi untuk melakukan automatis uppercase
	3. membuat output enkripsi ke sebuah file
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var pesan string = "AUSHAFY"
	var key string = "N"
	var newKey string = newKey(pesan, key)
	fmt.Println("Plaintext:", pesan)
	fmt.Println("Key:", key)
	fmt.Println("NewKey:", newKey)
	fmt.Println("Encrypt:", autoEncryption(pesan, newKey))
	fmt.Println("Decrypt:", autoDecryption(autoEncryption(pesan, newKey), newKey))
}

func newKey(pesan, key string) string {
	var newKey = strings.Replace(pesan, string(pesan[0]), key, 1)
	return newKey
}

func autoEncryption(pesan, key string) string {
	var pesanLen int = len(pesan)
	var newKey string = newKey(pesan, key)

	var pesanEncrypt []string
	var pesanEncryptString string
	var i int = 0
	for i < pesanLen {
		var first = pesan[i]
		var second = newKey[i]
		var total = (first + second) % 26
		total = total + 65
		pesanEncrypt = append(pesanEncrypt, string(total))
		i = i + 1
	}
	pesanEncryptString = strings.Join(pesanEncrypt, "")
	return pesanEncryptString
}

func autoDecryption(cipher, key string) string {
	var pesanDecryptSlices []string
	var pesanDecryptString string
	var i = 0
	for i < len(key) {
		var first = cipher[i]
		var second = key[i]
		var total = (first - second + 26) % 26
		total = total + 65
		pesanDecryptSlices = append(pesanDecryptSlices, string(total))
		i = i + 1
	}
	pesanDecryptString = strings.Join(pesanDecryptSlices, "")
	return pesanDecryptString
}
