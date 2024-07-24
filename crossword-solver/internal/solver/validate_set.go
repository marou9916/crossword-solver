package solver

func IsValidSetOfWords(setOfWords []string) bool {
	checker := make(map[string]struct{})

	for _, word := range setOfWords {
		if _, exists := checker[word]; exists {
			return false
		}

		checker[word] = struct{}{}
	}

	return true
}
