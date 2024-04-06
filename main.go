package main

import (
	"learngo-pockets/gordle/gordle"
	"os"
)

const maxAttempts = 6

func main() {

	solution := "hello"
	g := gordle.New(os.Stdin, solution, maxAttempts)
	g.Play()
}
