package main

import "strings"

func parseArgs(input string) []string {
	input = strings.TrimLeft(input, " \t\n")
	if len(input) == 0 {
		return nil
	}

	var current strings.Builder
	var result []string

	var (
		isSingle  bool
		isDouble  bool
		isEscaped bool
	)

	for _, ch := range input {
		switch {
		case ch == '\\' && !isSingle:
			if isEscaped {
				current.WriteRune(ch)
				isEscaped = false
			} else {
				isEscaped = true
			}
		case ch == '\'' && !isDouble && !isEscaped:
			isSingle = !isSingle
		case ch == '"' && !isSingle && !isEscaped:
			isDouble = !isDouble
		case (ch == ' ' || ch == '\t' || ch == '\n') && !isSingle && !isDouble && !isEscaped:
			if current.Len() > 0 {
				result = append(result, current.String())
				current.Reset()
			}
		default:
			if isEscaped && isDouble && ch != '"' && ch != '\\' {
				current.WriteRune('\\')
			}
			current.WriteRune(ch)
			isEscaped = false
		}
	}

	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}
