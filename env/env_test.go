package env

import (
	"os/user"
	"testing"
)

func Test_AddEnvVar(t *testing.T) {
	env := InitEnv()
	//Ensure that when we try to add a variable that already exists
	//we set it instead of adding another variable of the same name
	env.AddEnvVar("home", "two")
	testVartwo := env.GetEnvVar("home")
	if testVartwo != "two" {
		t.Error("Failed to set existing environment variable in AddEnvVar")
	}
}

func Test_SetEnvVar(t *testing.T) {
	env := InitEnv()
	//Ensure that if we try to set a non-existant variable, we add it instead
	env.SetEnvVar("asdf", "asdf")
	if env.CheckExists("asdf") != true {
		t.Error("Failed to add new variable when trying to set non-existing variable")
	}
}

func Test_GetEnvVar(t *testing.T) {
	env := InitEnv()
	nilVar := env.GetEnvVar("fdsa")
	if nilVar != "" {
		t.Error("Failed to return \"\" when variable doesn't exist")
	}
}

func Test_DeleteEnvVar(t *testing.T) {
	env := InitEnv()

	//Ensure that we can delete existing environment variables
	env.AddEnvVar("test", "test")
	env.DeleteEnvVar("test")
	if env.CheckExists("test") == true {
		t.Error("Failed to delete environment variable")
	}

	//Ensure that we get false when trying to delete non-existant environment variables
	ret := env.DeleteEnvVar("test")
	if ret != false {
		t.Error("Expecte false when trying to delete non-existant variable")
	}
}

func Test_InitEnv(t *testing.T) {
	env := InitEnv()

	usr, _ := user.Current()
	var strTests = map[string]string{
		"prompt": "$ ",
		"home":   usr.HomeDir,
		"editor": "mvim",
		"term":   "rxvt-unicode-256color",
		"pager":  "less",
		"user":   usr.Username,
	}

	//Check that our default environment variables exist
	for input, expect := range strTests {
		got := env.GetEnvVar(input)
		if expect != got {
			t.Error(input, ": Expected:", expect, "Got:", got)
		}
	}

	//Ensure the environment isn't dirty by default
	if env.Dirty != false {
		t.Error("env.Dirty != false")
	}

	//Ensure we have the right amount of variables
	if env.Count != len(strTests) {
		t.Error("Invalid variable count:", env.Count)
	}

	env.Print()
}
