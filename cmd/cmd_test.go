package cmd

import (
	"testing"
)

var testCmdStr string = "ls -hal"

func Test_GetNameStr(t *testing.T) {
	cmd := ParseInput(&testCmdStr)

	//Ensure that the name of the command, in this
	//case 'ls', is correct
	got := cmd.GetNameStr()
	if *got != "ls" {
		t.Error("Invalid name string from cmd.GetNameStr()")
	}
}

func Test_GetRawStr(t *testing.T) {
	cmd := ParseInput(&testCmdStr)

	if cmd.GetRawStr() != &testCmdStr {
		t.Error("Invalid raw string from cmd.GetRawStr()")
	}
}

func Test_GetElements(t *testing.T) {
	var expect int = 2
	cmd := ParseInput(&testCmdStr)

	if cmd.GetElements() != expect {
		t.Error("Invalid element count from cmd.GetElements()")
	}
}

func Test_GetTokens(t *testing.T) {
	expect := []string{"ls", "-hal"}
	cmd := ParseInput(&testCmdStr)

	arr := cmd.GetTokens()

	if len(arr) != len(expect) {
		t.Error("Invalid token array length, expect", len(expect),
			" Got", len(arr))
	}

	for i, got := range arr {
		if expect[i] != got {
			t.Error("Invalid token in token array from cmd.GetTokens(), expected",
				expect, " Got", got)
		}
	}
}
