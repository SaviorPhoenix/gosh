package main

import (
	"fmt"
	"github.com/sbinet/go-readline/pkg/readline"
	"os"
	"os/exec"
	"strings"

	"github.com/tywkeene/gosh/builtins"
	"github.com/tywkeene/gosh/cmd"
	"github.com/tywkeene/gosh/sh"
)

func executeCommand(c cmd.GoshCmd) error {
	str := strings.Join(c.Tokens[1:len(c.Tokens)], " ")
	parts := strings.Fields(str)
	file := c.GetNameStr()
	args := parts[0:len(parts)]

	run := exec.Command(file, args...)

	run.Stdout = os.Stdout
	run.Stdin = os.Stdin
	run.Stderr = os.Stderr

	err := run.Run()
	return err
}

func main() {
	shell.Sh.InitShell()
	env := shell.Sh.GetEnv()

	for {
		prompt := env.GetEnvVar("prompt")
		input := readline.ReadLine(&prompt)

		if *input == "" {
			continue
		}

		if env.VarCmp("history", "on") == true {
			readline.AddHistory(*input)
		}

		c := cmd.ParseInput(*input)
		builtin := builtins.CheckBuiltin(c)
		if builtin != nil {
			if err := builtin(c); err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			if err := executeCommand(c); err != nil {
				fmt.Println(err)
			}
		}
	}
}
