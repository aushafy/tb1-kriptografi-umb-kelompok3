package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var pesan, keyword, key string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan pesan: ")
	pesan, _ = reader.ReadString('\n')

	fmt.Print("Masukkan keyword: ")
	keyword, _ = reader.ReadString('\n')

	// remove all numeric, special character, and spacy from a string
	reg, err := regexp.Compile("[^a-zA-Z]+")
	proceedString := reg.ReplaceAllString(strings.Replace(pesan, " ", "", -1), "")
	proceedKeyword := reg.ReplaceAllString(strings.Replace(keyword, " ", "", -1), "")
	if err != nil {
		log.Fatal(err)
	}

	key = generateKey(proceedString, proceedKeyword)

	fmt.Println("pesan anda telah dienkrip menjadi:", cipherText(proceedString, key))
	fmt.Println("pesan cipher anda telah didekrip menjadi:", originalText(cipherText(proceedString, key), key))
}

func generateKey(str, key string) string {
	slicesKey := stringToSlices(key)
	lenStr := len(str)
	lenKey := len(slicesKey)
	substraction := lenStr - lenKey
	returnKey := ""
	// if the length of str and key are same, return key argument as string
	if lenStr == lenKey {
		returnKey = key
	} else { // if the length of str and key are not same, repeat the key until the length are same
		i := 0
		for i < substraction {
			slicesKey = append(slicesKey, slicesKey[i])
			i = i + 1
		}
		returnKey = strings.Join(slicesKey, "")
	}
	return returnKey
}

// String to Slice function in Golang. example: "AUSHAFY" would be convert to ["A", "U", "S", "H", "A", "F", "Y"]
func stringToSlices(str string) []string {
	var slicesString []string
	len := 0
	for _, letter := range str {
		len++
		slicesString = append(slicesString, string(letter))
	}
	return slicesString
}

func cipherText(str, key string) string {
	var cipherSlices []string
	var returnCipher string

	for i := range str {
		x := (str[i] + key[i]) % 26
		x = x + 65
		// fmt.Println(string(x))
		cipherSlices = append(cipherSlices, string(x))
	}
	returnCipher = strings.Join(cipherSlices, "")
	return returnCipher
}

func originalText(cipher, key string) string {
	var originalSlices []string
	var returnOriginal string
	for i := range cipher {
		x := (cipher[i] - key[i] + 26) % 26
		x = x + 65
		originalSlices = append(originalSlices, string(x))
	}
	returnOriginal = strings.Join(originalSlices, "")
	return returnOriginal
}
