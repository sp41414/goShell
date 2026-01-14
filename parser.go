package main

import "strings"

func parseArgs(input string) []string {
	var current strings.Builder
	result := []string{}
	isSingle := false
	isDouble := false

	for _, ch := range input {
		switch {
		case ch == '\'' && !isDouble:
			isSingle = !isSingle
		case ch == '"' && !isSingle:
			isDouble = !isDouble
		case (ch == ' ' || ch == '\t') && !isSingle && !isDouble:
			if current.Len() > 0 {
				result = append(result, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(ch)
		}
	}

	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}
