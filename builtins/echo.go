package builtins

import (
	"fmt"
	"strings"
)

func Echo(args []string) error {
	fmt.Print(strings.Join(args, " ") + "\n")

	return nil
}
