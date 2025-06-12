// Copyright 2024 Kenneth A. Giusti
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var debugGuess int

type Game struct {
	reader   *bufio.Reader
	writer   io.Writer
	secret   int
	attempts int
}

func NewGame(secret int, reader io.Reader, writer io.Writer) *Game {
	return &Game{
		reader:   bufio.NewReader(reader),
		writer:   writer,
		secret:   secret,
		attempts: 3,
	}
}

func (g *Game) Play() bool {
	for i := 0; i < g.attempts; i++ {
		fmt.Fprint(g.writer, "Enter your guess (1-100): ")
		input, err := g.reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(g.writer, "Error reading input. Please try again.")
			continue
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil || guess < 1 || guess > 100 {
			fmt.Fprintln(g.writer, "Invalid input. Please enter a number between 1 and 100.")
			continue
		}
		if guess == g.secret {
			fmt.Fprintln(g.writer, "You Won!")
			return true
		} else if i < g.attempts-1 {
			fmt.Fprintln(g.writer, "Sorry that did not match")
			if guess < g.secret {
				fmt.Fprintln(g.writer, "The mistery number is higher than your guess")
			} else {
				fmt.Fprintln(g.writer, "The mistery number is lower than your guess")
			}
			continue
		}
	}
	fmt.Fprintf(g.writer, "Sorry you lose, the correct number is %d!\n", g.secret)
	return false
}

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "guessinator",
		Short: "A number guessing game",
		RunE: func(cmd *cobra.Command, args []string) error {
			var secret int

			// Check if the debug-guess flag was provided
			if cmd.Flags().Changed("debug-guess") {
				// Perform the range check if the flag was provided
				if debugGuess < 1 || debugGuess > 100 {
					// Print the error message to stderr and return the error immediately
					return errors.New("Error: --debug-guess value must be between 1 and 100.")
				}
				// If the flag was provided and the value is valid, use it
				secret = debugGuess
			} else {
				// If the flag was not provided, generate a random number
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				secret = r.Intn(100) + 1
			}

			// ONLY create and play the game if the flag was valid or not provided
			game := NewGame(secret, os.Stdin, os.Stdout)
			game.Play()
			return nil
		},
	}

	rootCmd.Flags().IntVar(&debugGuess, "debug-guess", 0, "Set the secret number for testing (1-100)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
