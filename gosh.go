package main

import (
	"fmt"
	"github.com/sbinet/go-readline/pkg/readline"
	"os"
	"os/exec"
	"strings"

	"github.com/SaviorPhoenix/gosh/builtins"
	"github.com/SaviorPhoenix/gosh/cmd"
	"github.com/SaviorPhoenix/gosh/sh"
)

func executeCommand(c cmd.GoshCmd) int {
	str := strings.Join(c.Tokens[1:len(c.Tokens)], " ")
	parts := strings.Fields(str)
	file := c.GetNameStr()
	args := parts[0:len(parts)]

	run := exec.Command(*file, args...)

	run.Stdout = os.Stdout
	run.Stdin = os.Stdin
	run.Stderr = os.Stderr

	err := run.Run()
	if err != nil {
		fmt.Println(err)
	}
	return 0
}

func main() {
	shell.Sh.InitShell()
	env := shell.Sh.GetEnv()

	for {
		prompt := env.GetEnvVar("prompt")
		input := readline.ReadLine(&prompt)

		//Don't bother parsing input if it's empty
		if *input == "" {
			continue
		}

		if env.VarCmp("history", "on") == true {
			readline.AddHistory(*input)
		}

		c := cmd.ParseInput(input)
		//CheckBuiltin will return 1 if the command was a builtin,
		//So of course we want to skip executeCommand if it does
		if builtins.CheckBuiltin(c) == 1 {
			continue
		}

		executeCommand(c)
	}
}
