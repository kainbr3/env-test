package main

import (
	"fmt"
	"os"
)

func main() {
	CheckEnv()
}

func CheckEnv() {
	teste := os.Getenv("TEXTO")

	if teste == "ola mundo" {
		fmt.Println("deu bom")
	} else {
		fmt.Println("deu ruim")
	}
}
