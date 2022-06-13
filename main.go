package main

import (
	"fmt"
	"os"
)

func main() {
	teste := os.Getenv("TEXTO")

	if teste != "" {
		fmt.Println(teste)
	} else {
		fmt.Println("ola mundo")
	}
}
