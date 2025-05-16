# Guessinator

A simple number guessing game written in Go. The game generates a random number between 1 and 100, and the player has three attempts to guess it correctly.

## Features

- Three attempts to guess the correct number
- Helpful hints indicating if your guess is too high or too low
- Input validation
- Debug mode for testing (using --debug-guess flag)

## Building and Running

To build and run the game:

```bash
go build
./guessinator
```

To run with a specific number (debug mode):
```bash
./guessinator --debug-guess 42
```

## Testing

To run the tests:
```bash
go test -v
```

## License

This project is licensed under the [Apache License 2.0](LICENSE). 