//go:build unit
// +build unit

package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEnv(t *testing.T) {
	//os.Setenv("TEXTO", "ola mundo")
	teste := os.Getenv("TEXTO")

	fmt.Println("teste")

	if teste == "ola mundo" {
		fmt.Println("deu bom")
	} else {
		fmt.Println("deu ruim")
	}

	assert.NotEmpty(t, teste)
	assert.Equal(t, teste, "ola mundo")

}
