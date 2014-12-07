package builtins

import (
	"fmt"
	"github.com/SaviorPhoenix/gosh/cmd"
	"github.com/SaviorPhoenix/gosh/sh"
	"os"
)

type builtinFunc func(c cmd.GoshCmd) int

var builtins = map[string]builtinFunc{
	"exit": builtinFunc(
		func(c cmd.GoshCmd) int {
			fmt.Println("exit")
			os.Exit(0)

			//Not reached, obviously
			return 1
		}),

	"cd": builtinFunc(
		func(c cmd.GoshCmd) int {
			name := c.GetNameStr()
			if c.GetElements() == 1 && *name == "cd" {
				env := shell.Sh.GetEnv()
				os.Chdir(env.GetEnvVar("home"))
				return 1
			} else {
				tokens := c.GetTokens()
				os.Chdir(tokens[1])
			}
			return 1
		}),

	"add-var": builtinFunc(
		func(c cmd.GoshCmd) int {
			if c.GetElements() != 3 {
				fmt.Println("Usage: add-var <variable name> <variable value>")
			} else {
				env := shell.Sh.GetEnv()
				tokens := c.GetTokens()
				env.AddEnvVar(tokens[1], tokens[2])
			}
			return 1
		}),

	"print-var": builtinFunc(
		func(c cmd.GoshCmd) int {
			if c.GetElements() != 2 {
				fmt.Println("Usage: print-var <variable name>")
			} else {
				env := shell.Sh.GetEnv()
				tokens := c.GetTokens()
				printVar := env.GetEnvVar(tokens[1])
				fmt.Println(printVar)
			}
			return 1
		}),

	"delete-var": builtinFunc(
		func(c cmd.GoshCmd) int {
			if c.GetElements() != 2 {
				fmt.Println("Usage: delete-var <variable name>")
			} else {
				env := shell.Sh.GetEnv()
				tokens := c.GetTokens()
				delVar := env.GetEnvVar(tokens[1])
				env.DeleteEnvVar(delVar)
			}
			return 1
		}),
}

func CheckBuiltin(c cmd.GoshCmd) int {
	do := builtins[*c.NameStr]
	if do != nil {
		return do(c)
	} else {
		return 0
	}
}
