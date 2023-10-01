package string

import (
	"strings"
)

// ToLower converts a string to lowercase.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper converts a string to uppercase.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Concat concatenates two strings.
func Concat(s1, s2 string) string {
	return s1 + s2
}

// Reverse reverses a string.
func Reverse(s string) string {
	// Convert the string to a rune slice to handle Unicode characters correctly
	runes := []rune(s)
	length := len(runes)
	reversed := make([]rune, length)

	for i, j := length-1, 0; i >= 0; i, j = i-1, j+1 {
		reversed[j] = runes[i]
	}

	return string(reversed)
}

// Contains checks if a string contains a substring.
func Contains(s, substring string) bool {
	return strings.Contains(s, substring)
}

// CountSubstring counts the number of occurrences of a substring in a string.
func CountSubstring(s, substring string) int {
	return strings.Count(s, substring)
}

// ReplaceAll replaces all occurrences of a substring with another string.
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// TrimWhitespace trims leading and trailing whitespace from a string.
func TrimWhitespace(s string) string {
	return strings.TrimSpace(s)
}

// Split splits a string into substrings based on a delimiter and returns a slice of substrings.
func Split(s, delimiter string) []string {
	return strings.Split(s, delimiter)
}

// Join concatenates elements of a string slice into a single string using a separator.
func Join(strs []string, separator string) string {
	return strings.Join(strs, separator)
}

// Repeat returns a new string consisting of count copies of the input string.
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// TrimLeft trims leading characters from a string that match a specified set of characters.
func TrimLeft(s, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

// TrimRight trims trailing characters from a string that match a specified set of characters.
func TrimRight(s, cutset string) string {
	return strings.TrimRight(s, cutset)
}

// Fields splits a string into substrings at each instance of one or more consecutive white space characters.
func Fields(s string) []string {
	return strings.Fields(s)
}

// HasPrefix checks if a string starts with a specified prefix.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix checks if a string ends with a specified suffix.
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Map applies a mapping function to each Unicode code point in a string and returns the resulting slice of runes.
func Map(mapping func(rune) rune, s string) string {
	return strings.Map(mapping, s)
}

// Replace replaces all occurrences of a set of characters with a specified replacement.
func Replace(s string, oldnew ...string) string {
	return strings.NewReplacer(oldnew...).Replace(s)
}

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if s == t, -1 if s < t, and +1 if s > t.
func Compare(s, t string) int {
	return strings.Compare(s, t)
}

// ToValidUTF8 returns a copy of the string with invalid UTF-8 replaced by the replacement string.
func ToValidUTF8(s, replacement string) string {
	return strings.ToValidUTF8(s, replacement)
}

// TrimPrefix trims a specified prefix from a string.
func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// TrimSuffix trims a specified suffix from a string.
func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

// Count counts the number of non-overlapping instances of substr in s.
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int {
	return strings.Index(s, substr)
}

// IndexAny returns the index of the first instance of any Unicode code point from chars in s, or -1 if none are present.
func IndexAny(s, chars string) int {
	return strings.IndexAny(s, chars)
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// LastIndexAny returns the index of the last instance of any Unicode code point from chars in s, or -1 if none are present.
func LastIndexAny(s, chars string) int {
	return strings.LastIndexAny(s, chars)
}
