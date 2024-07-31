package solver

import (
    "crossword-solver/internal/model"
    "testing"
)

func TestCanPlaceWordVertically(t *testing.T) {
    grid := [][]rune{
        {'0', '0', '0', '0'},
        {'0', '.', '.', '0'},
        {'0', '0', '0', '0'},
        {'0', '.', '.', '0'},
    }

    tests := []struct {
        word     string
        pos      model.Position
        expected bool
    }{
        {"word", model.Position{Row: 0, Col: 0}, true},
        {"wor", model.Position{Row: 1, Col: 1}, false},
        {"wor", model.Position{Row: 1, Col: 2}, false},
        {"w", model.Position{Row: 1, Col: 3}, true},
        // Ajout de tests pour les lettres correspondantes
        {"anta", model.Position{Row: 0, Col: 2}, false},
    }

    for _, test := range tests {
        result := CanPlaceWordVertically(grid, test.word, test.pos)
        if result != test.expected {
            t.Errorf("CanPlaceWordVertically(%q, %v) = %v; want %v", test.word, test.pos, result, test.expected)
        }
    }
}
