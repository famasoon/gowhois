package main

import (
	"fmt"
	"log"
	"os"

	"github.com/famasoon/gowhois/whois"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s domain\n", os.Args[0])
		fmt.Printf("Example: %s example.com\n", os.Args[0])
		fmt.Printf("Example: %s 1.1.1.1\n", os.Args[0])
		os.Exit(0)
	}

	result, err := whois.Whois(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Get whois: " + result)
}
