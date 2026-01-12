package builtins

import "fmt"

type Command map[string]func(args []string) error

var Commands = Command{
	"exit": Exit,
	"echo": Echo,
}

func FindCommand(name string) (func(args []string) error, error) {
	if cmdFunc, ok := Commands[name]; ok {
		return cmdFunc, nil
	}
	return nil, fmt.Errorf("%s: command not found\n", name)
}
