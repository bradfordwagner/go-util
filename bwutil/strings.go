package bwutil

import "strings"

// StringsSplitAndTrim - splits a string by a separator and trims the parts
func StringsSplitAndTrim(s, sep string) []string {
	parts := strings.Split(s, sep)
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}
