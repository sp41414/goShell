package builtins

import (
	"os"
	"strconv"
)

func Exit(args []string) error {
	if len(args) == 0 {
		os.Exit(0)
	}

	statusCode, err := strconv.Atoi(args[0])
	if err != nil {
		os.Exit(2)
	}

	os.Exit(statusCode)
	return nil
}
