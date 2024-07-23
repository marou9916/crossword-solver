#!/usr/bin/bash

# Créer la structure des répertoires
mkdir -p crossword-solver/cmd/crosswordsolver
mkdir -p crossword-solver/internal/solver
mkdir -p crossword-solver/internal/model
mkdir -p crossword-solver/test

# Créer les fichiers .go avec du contenu initial
echo 'package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
}
' > crossword-solver/cmd/crosswordsolver/main.go

echo 'package solver

// Example function in backtrack
func Backtrack() {
    // Function logic here
}
' > crossword-solver/internal/solver/backtrack.go

echo 'package solver

// Example function in validate
func Validate() {
    // Function logic here
}
' > crossword-solver/internal/solver/validate.go

echo 'package solver

// Example function in place
func Place() {
    // Function logic here
}
' > crossword-solver/internal/solver/place.go

echo 'package solver

// Example function in positions
func Positions() {
    // Function logic here
}
' > crossword-solver/internal/solver/positions.go

echo 'package model

// Example function or struct in position
type Position struct {
    X, Y int
}
' > crossword-solver/internal/model/position.go

echo 'package test

import "testing"

func TestExample(t *testing.T) {
    // Example test
    if 1+1 != 2 {
        t.Errorf("1+1 should equal 2")
    }
}
' > crossword-solver/test/solver_test.go

