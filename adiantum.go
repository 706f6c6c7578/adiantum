package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"log"

	"lukechampine.com/adiantum"
)

func encryptFromStdin() {
	input, _ := ioutil.ReadAll(os.Stdin)
	if len(input) == 0 {
		return
	}
	noNewline := input[len(input)-1] != '\n'
	input = bytes.TrimRight(input, "\n")
	if len(input) < 16 {
		zeros := bytes.Repeat([]byte{0}, 16-len(input))
		input = append(input, zeros...)
	}
	secretkey, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Println("Error reading the key file: ", err)
		os.Exit(1)
	}
	nonce, err := ioutil.ReadFile(flag.Arg(1))
	if err != nil {
		fmt.Println("Error reading the nonce file: ", err)
		os.Exit(1)
	}
	secretkey = bytes.TrimRight(secretkey, "\r\n")
	nonce = bytes.TrimRight(nonce, "\r\n")
	keyInput, err := hex.DecodeString(string(secretkey))
	if err != nil {
		fmt.Printf("Error decoding the key: %v\n", err)
		os.Exit(1)
	}
	if len(keyInput) != 32 {
		log.Fatalf("Your key is %d hex bytes in size. Exactly 32 hex bytes are required.\n", len(keyInput))
		os.Exit(1)
	}
	key := keyInput
	tweak := nonce // can be any length, but should be at least 12 bytes.
	cipher := adiantum.New(key)
	output := cipher.Encrypt(input, tweak)
	if noNewline {
		fmt.Print(string(output))
	} else {
		fmt.Println(string(output))
	}
}

func decryptFromStdin() {
	input, _ := ioutil.ReadAll(os.Stdin)
	if len(input) == 0 {
		return
	}
	noNewline := input[len(input)-1] != '\n'
	input = bytes.TrimRight(input, "\n")
	if len(input) < 16 {
		zeros := bytes.Repeat([]byte{0}, 16-len(input))
		input = append(input, zeros...)
	}
	secretkey, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Println("Error reading the key file: ", err)
		os.Exit(1)
	}
	nonce, err := ioutil.ReadFile(flag.Arg(1))
	if err != nil {
		fmt.Println("Error reading the nonce file: ", err)
		os.Exit(1)
	}
	secretkey = bytes.TrimRight(secretkey, "\r\n")
	nonce = bytes.TrimRight(nonce, "\r\n")
	keyInput, err := hex.DecodeString(string(secretkey))
	if err != nil {
		fmt.Printf("Error decoding the key: %v\n", err)
		os.Exit(1)
	}
	if len(keyInput) != 32 {
		log.Fatalf("Your key is %d hex bytes in size. Exactly 32 hex bytes without CRLF/LF are required.\n", len(keyInput))
		os.Exit(1)
	}
	key := keyInput
	tweak := nonce // can be any length, but should be at least 12 bytes.
	cipher := adiantum.New(key)
	output := cipher.Decrypt(input, tweak)
	if noNewline {
		fmt.Print(string(output))
	} else {
		fmt.Println(string(output))
	}
}

func main() {

	decryptFlag := flag.Bool("d", false, "Decrypt: adiantum -d keyfile noncefile < infile > outfile")

	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Please provide both a key file and a nonce file.")
		os.Exit(1)
	}

	if decryptFlag != nil && *decryptFlag {
		decryptFromStdin()
	} else {
		encryptFromStdin()
	}
}
