package app

import "fmt"

type App struct {
	Greeting string
}

func (a App) SayGreetings() {
	fmt.Println(a.Greeting)
}
