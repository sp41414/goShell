package builtins

import "fmt"

type Command map[string]func(args []string) error

var Commands Command

func InitCommands() Command {
	Commands = Command{
		"exit": Exit,
		"echo": Echo,
		"type": Type,
	}

	return Commands
}

func FindCommandCallback(name string, commands Command) (func(args []string) error, error) {
	if cmdFunc, ok := commands[name]; ok {
		return cmdFunc, nil
	}
	return nil, fmt.Errorf("%s: command not found\n", name)
}

func FindCommand(name string, commands Command) bool {
	if _, ok := commands[name]; ok {
		return true
	}
	return false
}
