package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const solutionLength = 5

// Game holds all the information we need to play a game of gordle
type Game struct {
	reader *bufio.Reader
}

// New returns a Game, which can be used to play!
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}

// Play runs the game
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")
	fmt.Printf("Enter a guess:\n")
}

// ask reads input until a valid suggestion is made & returned
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr,
				"Gordle failed to read your guess: %s\n", err.Error())
		}

		guess := []rune(string(playerInput))

		if len(guess) != solutionLength {
			_, _ = fmt.Fprintf(os.Stderr,
				"Your attempt is invalid with Gordle's solution!"+
					"Expected %d characters, got %d.\n", solutionLength, len(guess))
		} else {
			return guess
		}
	}
}
