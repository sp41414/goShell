package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		"pwd":  Pwd,
		"cd":   Cd,
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

func FindExecutable(paths []string, executable string) (string, bool) {
	for _, path := range paths {
		if path == "" {
			continue
		}
		fullPath := filepath.Join(path, executable)

		// Checks if path exists, is a regular file, has executable permissions.
		if fi, err := os.Stat(fullPath); err == nil && fi.Mode().IsRegular() && fi.Mode().Perm()&0111 != 0 {
			return fullPath, true
		}
	}
	return "", false
}

func Execute(file string, args []string) error {
	path, err := ParsePath()
	if err != nil {
		return err
	}
	fullPath, found := FindExecutable(path, file)

	if !found {
		if fi, err := os.Stat(file); err == nil && fi.Mode().IsRegular() && fi.Mode().Perm()&0111 != 0 {
			if !strings.ContainsRune(file, filepath.Separator) {
				fullPath = "./" + file
			} else {
				fullPath = file
			}
			found = true
		}
	}

	if !found {
		return fmt.Errorf("%s: command not found", file)
	}

	cmd := exec.Command(fullPath, args...)
	cmd.Args[0] = file
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return nil
		}
		return err
	}
	return nil
}
