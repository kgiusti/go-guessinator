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
	"bytes"
	"strings"
	"testing"
)

func TestGame(t *testing.T) {
	tests := []struct {
		name           string
		secret         int
		inputs         []string
		expectedOutput string
		expectedWin    bool
	}{
		{
			name:   "Win on first guess",
			secret: 42,
			inputs: []string{"42\n"},
			expectedOutput: `Enter your guess (1-100): You Won!
`,
			expectedWin: true,
		},
		{
			name:   "Win on second guess with higher hint",
			secret: 75,
			inputs: []string{"50\n", "75\n"},
			expectedOutput: `Enter your guess (1-100): Sorry that did not match
The mistery number is higher than your guess
Enter your guess (1-100): You Won!
`,
			expectedWin: true,
		},
		{
			name:   "Win on second guess with lower hint",
			secret: 25,
			inputs: []string{"50\n", "25\n"},
			expectedOutput: `Enter your guess (1-100): Sorry that did not match
The mistery number is lower than your guess
Enter your guess (1-100): You Won!
`,
			expectedWin: true,
		},
		{
			name:   "Lose after three wrong guesses",
			secret: 42,
			inputs: []string{"10\n", "20\n", "30\n"},
			expectedOutput: `Enter your guess (1-100): Sorry that did not match
The mistery number is higher than your guess
Enter your guess (1-100): Sorry that did not match
The mistery number is higher than your guess
Enter your guess (1-100): Sorry you lose, the correct number is 42!
`,
			expectedWin: false,
		},
		{
			name:   "Invalid input followed by win",
			secret: 42,
			inputs: []string{"invalid\n", "42\n"},
			expectedOutput: `Enter your guess (1-100): Invalid input. Please enter a number between 1 and 100.
Enter your guess (1-100): You Won!
`,
			expectedWin: true,
		},
		{
			name:   "Out of range input followed by win",
			secret: 42,
			inputs: []string{"101\n", "42\n"},
			expectedOutput: `Enter your guess (1-100): Invalid input. Please enter a number between 1 and 100.
Enter your guess (1-100): You Won!
`,
			expectedWin: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input buffer with test inputs
			input := strings.Join(tt.inputs, "")
			reader := strings.NewReader(input)

			// Create output buffer to capture game output
			var output bytes.Buffer

			// Create and play game
			game := NewGame(tt.secret, reader, &output)
			win := game.Play()

			// Check if output matches expected
			if output.String() != tt.expectedOutput {
				t.Errorf("expected output:\n%q\ngot:\n%q", tt.expectedOutput, output.String())
			}

			// Check if win/lose status matches expected
			if win != tt.expectedWin {
				t.Errorf("expected win=%v, got win=%v", tt.expectedWin, win)
			}
		})
	}
}
