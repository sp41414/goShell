package builtins

import (
	"fmt"
	"os"
	"path/filepath"
)

func Type(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("type: missing argument\n")
	}
	command := args[0]

	if _, ok := Commands[command]; ok {
		fmt.Printf("%s is a shell builtin\n", command)
		return nil
	}

	paths, err := ParsePath()
	if err != nil {
		return err
	}

	for _, path := range paths {
		if path == "" {
			continue
		}
		fullPath := filepath.Join(path, command)

		// checks if path exists, is a regular file, has executable permissions
		if fi, err := os.Stat(fullPath); err == nil && fi.Mode().IsRegular() && fi.Mode().Perm()&0111 != 0 {
			fmt.Printf("%s is %s\n", command, fullPath)
			return nil
		}
	}

	return fmt.Errorf("%s: not found", command)
}
