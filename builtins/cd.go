package builtins

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Cd(args []string) error {
	var path string
	home, err := os.UserHomeDir()
	if err != nil {
		home = "/"
	}

	if len(args) == 0 || args[0] == "" {
		path = home
	} else {
		path = args[0]
	}

	if strings.HasPrefix(path, "~") {
		trimmed := strings.TrimPrefix(path, "~")
		path = filepath.Join(home, trimmed)
	}

	if path == "-" {
		path = os.Getenv("OLDPWD")
		if path == "" {
			return fmt.Errorf("cd: OLDPWD is not set")
		}
	}

	cwd, _ := os.Getwd()
	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("cd: %s: No such file or directory", path)
	}

	os.Setenv("OLDPWD", cwd)

	return nil
}
