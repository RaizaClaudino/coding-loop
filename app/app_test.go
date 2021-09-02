package app_test

import (
	"testing"
	"wegift/app"

	"github.com/stretchr/testify/assert"
)

func TestNewApp_Sucess(t *testing.T) {
	newApp := app.App{Greeting: "Hello"}

	assert.Equal(t, newApp.Greeting, "Hello")
}
