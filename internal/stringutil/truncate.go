// Package stringutil provides shared string operations.
package stringutil

import "unicode/utf8"

// SafeTruncate cuts a string to maxBytes without breaking a UTF-8 character sequence.
// It assumes valid UTF-8 input and a nonnegative maxBytes. Callers add any suffix.
func SafeTruncate(s string, maxBytes int) string {
	if len(s) <= maxBytes {
		return s
	}
	// Step backward (at most 3 bytes) until we find the start of a valid UTF-8 character.
	for maxBytes > 0 && !utf8.RuneStart(s[maxBytes]) {
		maxBytes--
	}
	return s[:maxBytes]
}
