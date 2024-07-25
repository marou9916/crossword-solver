# Crossword Solver
This project is a renewed attempt to solve the Crossword Solver problem, which I had tackled a few months ago. At that time, I didn't have a good grasp of backtracking, which posed some difficulties. Determined to overcome this challenge, I studied and practiced backtracking through various exercises on GeeksForGeeks. You can find these solutions in my repository `problem-solving-practice`, under the 'backtracking' folder. With this experience, I now feel ready to apply my new skills to solve the Crossword Solver.

## Introduction

The Crossword Solver is a crossword puzzle solver written in Go. It takes a puzzle grid and a set of words as input and finds the positions where the words can be placed in the grid.

## Structure du Projet
```
crossword-solver/
├── cmd/
│ └── crosswordsolver/
│ └── main.go
├── internal/
│ ├── validator/
│ │ ├── validate_string.go
│ │ ├── validate_string_test.go
│ │ ├── validate_setofwords.go
│ │ └── validate_setofwords_test.go
│ ├── solver/
│ │ ├── backtrack.go
│ │ ├── place.go
│ │ ├── positions.go
│ │ └── solve.go
│ └── model/
│ └── position.go
├── test/
│ └── solver_test.go
├── go.mod
└── go.sum
```

## Features

- **Puzzle Format Validation**: Checks that the provided puzzle grid is in a valid format.
- **Word Set Validation**: Ensures the provided set of words contains no duplicates.
- **Word Placement**:  Determines possible positions to place words in the grid.

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


