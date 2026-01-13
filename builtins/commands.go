package builtins

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

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

func ParsePath() ([]string, error) {
	path, found := os.LookupEnv("PATH")
	if !found {
		return []string{}, fmt.Errorf("PATH environment variable not found\n")
	}

	var splitPath []string
	if runtime.GOOS == "windows" {
		splitPath = strings.Split(path, ";")
	} else {
		splitPath = strings.Split(path, ":")
	}

	return splitPath, nil
}
