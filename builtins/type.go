package builtins

import "fmt"

func Type(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("type: missing argument\n")
	}
	command := args[0]

	if _, ok := Commands[command]; ok {
		fmt.Printf("%s is a shell builtin\n", command)
		return nil
	}

	return fmt.Errorf("%s: not found", command)
}
