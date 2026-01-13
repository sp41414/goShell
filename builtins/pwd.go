package builtins

import (
	"fmt"
	"os"
)

func Pwd(args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", dir)
	return nil
}
