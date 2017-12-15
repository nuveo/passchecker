package gopsc

import (
	"strings"
)

var (
	enUS = map[rune]float64{
		'e': 0.87298,
		't': 0.90944,
		'a': 0.91833,
		'o': 0.92493,
		'i': 0.93034,
		'n': 0.93251,
		's': 0.93673,
		'h': 0.93906,
		'r': 0.94013,
		'd': 0.95747,
		'l': 0.95975,
		'c': 0.97218,
		'u': 0.97242,
		'm': 0.97594,
		'w': 0.97640,
		'f': 0.97772,
		'g': 0.97985,
		'y': 0.98026,
		'p': 0.98071,
		'b': 0.98508,
		'v': 0.99022,
		'k': 0.99228,
		'j': 0.99847,
		'x': 0.99850,
		'q': 0.99905,
		'z': 0.99926,
		'E': 0.87298,
		'T': 0.90944,
		'A': 0.91833,
		'O': 0.92493,
		'I': 0.93034,
		'N': 0.93251,
		'S': 0.93673,
		'H': 0.93906,
		'R': 0.94013,
		'D': 0.95747,
		'L': 0.95975,
		'C': 0.97218,
		'U': 0.97242,
		'M': 0.97594,
		'W': 0.97640,
		'F': 0.97772,
		'G': 0.97985,
		'Y': 0.98026,
		'P': 0.98071,
		'B': 0.98508,
		'V': 0.99022,
		'K': 0.99228,
		'J': 0.99847,
		'X': 0.99850,
		'Q': 0.99905,
		'Z': 0.99926,
		'1': 0.98700,
		'2': 0.98600,
		'3': 0.98500,
		'4': 0.98400,
		'5': 0.98300,
		'6': 0.98200,
		'7': 0.98100,
		'8': 0.98000,
		'9': 0.97900,
		'0': 0.97800,
	}

	enUSSequence = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

// CheckStrength returns a value that represents
// the strength of a password, the greater the
// number the greater the strength of the password
func CheckStrength(password string) (strength float64) {

	sumRuneOccurrences := make(map[rune]float64)

	for i, c := range password {
		//runeValue, width := utf8.DecodeRuneInString(c)
		//plen += width

		// default value
		charValue := 1.0

		// get rune value by frequency in language
		if val, ok := enUS[c]; ok {
			charValue = val
		}

		// penalty for char sequencing
		if i > 0 {
			if strings.Contains(enUSSequence, password[0:i]) {
				charValue = charValue - (charValue / 4.0)
			}
		}

		// registry number of rune occurrence
		sumRuneOccurrences[c]++

		// compute rune value
		strength += charValue / sumRuneOccurrences[c]
	}

	return
}
