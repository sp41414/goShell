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

	orgStdout := os.Stdout
	orgStderr := os.Stderr

	for {
		fmt.Fprintf(orgStdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalf("Error: could not read command %v", err)
		}
		command = strings.TrimSpace(command)
		args := parseArgs(command)
		if len(args) <= 0 {
			continue
		}

		args, stdoutFile, stderrFile, appendOut, appendErr := parseRedirect(args)

		var stdoutFileCurr *os.File
		var stderrFileCurr *os.File

		if stdoutFile != "" {
			flags := os.O_WRONLY | os.O_CREATE
			if appendOut {
				flags |= os.O_APPEND
			} else {
				flags |= os.O_TRUNC
			}
			f, err := os.OpenFile(stdoutFile, flags, 0644)
			if err != nil {
				fmt.Fprintln(orgStderr, err)
				continue
			}
			os.Stdout = f
			stdoutFileCurr = f
		}

		if stderrFile != "" {
			flags := os.O_WRONLY | os.O_CREATE
			if appendErr {
				flags |= os.O_APPEND
			} else {
				flags |= os.O_TRUNC
			}
			f, err := os.OpenFile(stderrFile, flags, 0644)
			if err != nil {
				fmt.Fprintln(orgStderr, err)
				continue
			}
			os.Stderr = f
			stderrFileCurr = f
		}

		callback, err := builtins.FindCommandCallback(args[0], commands)
		if err != nil {
			execErr := builtins.Execute(args[0], args[1:])
			if execErr != nil {
				fmt.Println(execErr)
			}
			goto cleanup
		}

		err = callback(args[1:])
		if err != nil {
			fmt.Println(err)
		}

	cleanup:
		if stdoutFileCurr != nil {
			stdoutFileCurr.Sync()
			stdoutFileCurr.Close()
		}
		if stderrFileCurr != nil {
			stderrFileCurr.Sync()
			stderrFileCurr.Close()
		}
		os.Stdout = orgStdout
		os.Stderr = orgStderr
	}
}
