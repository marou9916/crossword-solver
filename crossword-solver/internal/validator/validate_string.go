package validator

import (
	"regexp"
	"strings"
)

// fonction de vérification de la validité du format de la chaîne de caractères représentant le puzzle
func IsValidString(puzzle string) bool {
	formatOfValidPattern := `^[012\. \n]*[012]$`

	matched, err := regexp.MatchString(formatOfValidPattern, puzzle)
	if err != nil || !matched {
		return false
	}

	lines := strings.Split(puzzle, "\n")
	if len(lines) < 2 {
		return false
	}
	lineLength := len(lines[0])
	for _, line := range lines {
		if len(line) != lineLength {
			return false
		}
	}

	return true
}

