package main

import "fmt"

func main() {
	greeting := "Hi"
	router("pong", handler{greeting: greeting}.printString)
}

type handler struct {
	greeting string
}

func (h.handler) printString(s string) {
	fmt.Println(h.greeting, s)
}

func router(s string, fn func(string)) {
	fn(s)
}
