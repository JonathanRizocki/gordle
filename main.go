package main

import (
	"fmt"
	"learngo-pockets/gordle/gordle"
	"os"
)

const maxAttempts = 6

func main() {
	corpus, err := gordle.ReadCorpus("corpus/english.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
		return
	}

	// Create the game
	g, err := gordle.New(corpus, gordle.WithMaxAttempts(maxAttempts))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start game :%s", err)
		return
	}
	g.Play()
}
