package env

import (
	"os"
	"os/user"
	"testing"
)

//Ensure that when we try to add a variable that already exists
//we set it instead of adding another variable of the same name
func Test_AddEnvVar(t *testing.T) {
	env := InitEnv()
	env.AddEnvVar("home", "two")
	testVartwo := env.GetEnvVar("home")
	if testVartwo != "two" {
		t.Error("Failed to set existing environment variable in AddEnvVar")
	}
}

//Ensure that if we try to set a non-existant variable, we add it instead
func Test_SetEnvVar(t *testing.T) {
	env := InitEnv()
	env.SetEnvVar("asdf", "asdf")
	if env.CheckExists("asdf") != true {
		t.Error("Failed to add new variable when trying to set non-existing variable")
	}
}

//Ensure that we get "" if we try to env.GetEnvVar a variable that doesn't exist
func Test_GetEnvVar(t *testing.T) {
	env := InitEnv()
	nilVar := env.GetEnvVar("fdsa")
	if nilVar != "" {
		t.Error("Failed to return \"\" when variable doesn't exist")
	}
}

//Ensure that we can delete existing variables and get false when we try to delete non
//existant variables
func Test_DeleteEnvVar(t *testing.T) {
	env := InitEnv()
	env.AddEnvVar("test", "test")
	env.DeleteEnvVar("test")

	//Ensure that we can delete existing environment variables
	if env.CheckExists("test") == true {
		t.Error("Failed to delete environment variable")
	}

	//Ensure that we get false when trying to delete non-existant environment variables
	if env.DeleteEnvVar("test") != false {
		t.Error("Expecte false when trying to delete non-existant variable")
	}
}

//Ensure that we get a correct default variable when calling InitEnv()
func Test_InitEnv(t *testing.T) {
	env := InitEnv()

	usr, _ := user.Current()
	pwd, _ := os.Getwd()

	//These are the default environment variables that come
	//pre loaded and hardcoded. The user can still set them
	//with the 'set-var' builtin command.
	var strTests = map[string]string{
		"prompt":       "$ ",
		"pwd":          pwd,
		"history":      "on",
		"history_file": usr.HomeDir + "/.gosh_history",
		"home":         usr.HomeDir,
		"editor":       os.Getenv("EDITOR"),
		"term":         os.Getenv("TERM"),
		"pager":        os.Getenv("PAGER"),
		"user":         usr.Username,
	}

	//Check that our default environment variables exist
	for input, expect := range strTests {
		got := env.GetEnvVar(input)
		if expect != got {
			t.Error(input, ": Expected:", expect, "Got:", got)
		}
	}

	//Ensure the environment isn't dirty by default
	if env.dirty != false {
		t.Error("env.dirty != false")
	}

	//Ensure we have the right amount of variables
	if env.count != len(strTests) {
		t.Error("Invalid variable count:", env.count)
	}

	env.Print()
}
