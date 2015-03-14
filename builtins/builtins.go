package builtins

import (
	"fmt"
	"github.com/SaviorPhoenix/gosh/cmd"
	"github.com/SaviorPhoenix/gosh/sh"
	"os"
)

type builtinFunc func(c cmd.GoshCmd) error

var builtins = map[string]builtinFunc{
	"exit": builtinFunc(
		func(c cmd.GoshCmd) error {
			fmt.Println("exit")
			return func() error {
				os.Exit(0)
				return nil
			}()
		}),

	"cd": builtinFunc(
		func(c cmd.GoshCmd) error {
			if c.GetElements() == 1 {
				env := shell.Sh.GetEnv()
				os.Chdir(env.GetEnvVar("home"))
				return nil
			} else {
				tokens := c.GetTokens()
				os.Chdir(tokens[1])
			}
			return nil
		}),

	"add-var": builtinFunc(
		func(c cmd.GoshCmd) error {
			if c.GetElements() != 3 {
				fmt.Println("Usage: add-var <variable name> <variable value>")
			} else {
				env := shell.Sh.GetEnv()
				tokens := c.GetTokens()
				env.AddEnvVar(tokens[1], tokens[2])
			}
			return nil
		}),

	"set-var": builtinFunc(
		func(c cmd.GoshCmd) error {
			if c.GetElements() != 3 {
				fmt.Println("Usage: set-var <variable name> <new variable value>")
			} else {
				env := shell.Sh.GetEnv()
				tokens := c.GetTokens()
				env.SetEnvVar(tokens[1], tokens[2])
			}
			return nil
		}),

	"print-var": builtinFunc(
		func(c cmd.GoshCmd) error {
			if c.GetElements() != 2 {
				fmt.Println("Usage: print-var <variable name>")
			} else {
				env := shell.Sh.GetEnv()
				tokens := c.GetTokens()

				printVar := env.GetEnvVar(tokens[1])
				if printVar == "" {
					fmt.Println("No such variable:", tokens[1])
					return nil
				}

				fmt.Println(printVar)
			}
			return nil
		}),

	"delete-var": builtinFunc(
		func(c cmd.GoshCmd) error {
			if c.GetElements() != 2 {
				fmt.Println("Usage: delete-var <variable name>")
			} else {
				env := shell.Sh.GetEnv()
				tokens := c.GetTokens()
				delVar := env.GetEnvVar(tokens[1])

				if env.DeleteEnvVar(delVar) == false {
					fmt.Println("No such variable:", tokens[1])
				}
			}
			return nil
		}),
}

func CheckBuiltin(c cmd.GoshCmd) builtinFunc {
	return builtins[c.NameStr]
}
