# Crossword Solver
This project is a renewed attempt to solve the Crossword Solver problem, which I had tackled a few months ago. At that time, I didn't have a good grasp of backtracking, which posed some difficulties. Determined to overcome this challenge, I studied and practiced backtracking through various exercises on GeeksForGeeks. You can find these solutions in my repository `problem-solving-practice`, under the 'backtracking' folder. With this experience, I now feel ready to apply my new skills to solve the Crossword Solver.

## Introduction

The Crossword Solver is a crossword puzzle solver written in Go. It takes a puzzle grid and a set of words as input and finds the positions where the words can be placed in the grid.

## Project Structure
```
crossword-solver/
├── cmd/
│   └── crosswordsolver/
│       └── main.go
├── internal/
│   ├── model/
│   │   └── position.go
│   ├── placement/
│   │   ├── can_place_word_horizontally.go
│   │   ├── can_place_word_horizontally_test.go
│   │   ├── place_word_horizontally.go
│   │   ├── place_word_horizontally_test.go
│   │   ├── can_place_word_vertically.go
│   │   ├── can_place_word_vertically_test.go
│   │   └── place_word_vertically.go
│   │   ├── place_word_vertically_test.go
│   ├── remover/
│   │   ├── remove_word_horizontally.go
│   │   ├── remove_word_horizontally_test.go
│   │   ├── remove_word_vertically.go
│   │   └── remove_word_vertically_test.go
│   ├── solver/
│   │   ├── try_to_find_solution_with_backtracking.go
│   │   ├── try_to_find_solution_with_backtracking_test.go
│   │   ├── convert_puzzle_to_grid.go
│   │   ├── convert_puzzle_to_grid_test.go
│   │   ├── fill_grid_with_content.go
│   │   ├── fill_grid_with_content_test.go
│   │   ├── find_positions.go
│   │   ├── find_positions_test.go
│   │   ├── get_dot_positions.go
│   │   ├── get_dot_positions_test.go
│   │   └── solve.go
│   ├── printer/
│   │   └── print_grid.go
│   └── validator/
│       ├── validate_string.go
│       ├── validate_string_test.go
│       ├── validate_setofwords.go
│       └── validate_setofwords_test.go
├── setup_crossWordSolver.sh
├── go.mod
└── go.sum
```

## Features

**Puzzle Format Validation**: Checks that the provided puzzle grid is in a valid format.

**Word Set Validation**: Ensures the provided set of words contains no duplicates.

**Grid Conversion**: Converts the puzzle grid from a string format to a 2D rune slice.

**Dot Position Identification**: Identifies the positions in the grid where words can potentially be placed.

**Word Placement Checking**: Checks if a word can be placed at a specific position horizontally or vertically.

**Word Placement**: Places words in the grid if they can be placed without conflicts.

**Word Removal**: Removes words from the grid, useful for backtracking.

**Backtracking Solver**: Uses backtracking to try to place all words in the grid.

**Grid Printing**: Prints the current state of the grid.

## Utilisation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/marou9916/crossword-solver.git
    cd crossword-solver
    ```

2. **Install dependencies**:

    ```bash
    go mod tidy
    ```

3. **Run the tests**:

    ```bash
    go test ./...
    ```

## Contributions

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes you would like to make.

## Licence

This project is licensed under the MIT License. See the LICENSE file for details.


