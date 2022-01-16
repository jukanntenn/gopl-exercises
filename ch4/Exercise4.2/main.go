package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var algo = flag.String("algo", "sha256", "flag = sha256 | sha384 | sha512")

func main() {
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		bytes := input.Bytes()
		switch *algo {
		case "sha256":
			fmt.Printf("sha256: %x\n", sha256.Sum256(bytes))
		case "sha384":
			fmt.Printf("sha384: %x\n", sha512.Sum384(bytes))
		case "sha512":
			fmt.Printf("sha512: %x\n", sha512.Sum512(bytes))
		}
	}
}
