package utils

import "strings"

func CleanInput(str string) []string {

	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
