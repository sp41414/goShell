package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sp41414/goShell/builtins"
)

func main() {
	commands := builtins.InitCommands()

	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalf("Error: could not read command %v", err)
		}
		command = strings.TrimSpace(command)
		args := parseArgs(command)
		if len(args) <= 0 {
			continue
		}

		callback, err := builtins.FindCommandCallback(args[0], commands)
		if err != nil {
			execErr := builtins.Execute(args[0], args[1:])
			if execErr != nil {
				fmt.Println(execErr)
			}
			continue
		}

		err = callback(args[1:])
		if err != nil {
			fmt.Println(err)
		}
	}
}
