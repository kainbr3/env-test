package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEnv(t *testing.T) {
	//os.Setenv("TEXTO", "ola mundo")
	teste := os.Getenv("TEXTO")

	assert.NotEmpty(t, teste)
	assert.Equal(t, teste, "ola mundo")

}
