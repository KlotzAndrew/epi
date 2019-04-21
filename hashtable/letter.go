package hashtable

func IsLetterConstructable(letter, mag string) bool {
	freq := map[rune]int{}
	for _, r := range letter {
		freq[r]++
	}

	for _, r := range mag {
		switch count, ok := freq[r]; {
		case count == 1:
			delete(freq, r)
		case ok:
			freq[r]--
		}

		if len(freq) == 0 {
			return true
		}
	}

	return false
}
