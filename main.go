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
	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalf("Error: could not read command %v", err)
		}
		args := strings.Fields(command)
		if len(args) <= 0 {
			continue
		}

		callback, err := builtins.FindCommand(args[0])
		if err != nil {
			fmt.Print(err)
			continue
		}
		callback(args[1:])
	}
}
