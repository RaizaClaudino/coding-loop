package main

import "wegift/app"

func main() {
	newApp := app.App{Greeting: "Hello world!"}

	newApp.SayGreetings()
}
