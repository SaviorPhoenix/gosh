package cmd

import (
	"strings"
)

type GoshCmd struct {
	NameStr  string
	RawStr   string
	Elements int
	Tokens   []string
}

func (c GoshCmd) GetNameStr() string {
	return c.NameStr
}

func (c GoshCmd) GetRawStr() string {
	return c.RawStr
}

func (c GoshCmd) GetElements() int {
	return c.Elements
}

func (c GoshCmd) GetTokens() []string {
	return c.Tokens
}

func ParseInput(input string) GoshCmd {
	var ret GoshCmd

	tokens := strings.Split(input, " ")
	//Plus one since we start at 0, and we can't have null input
	size := strings.Count(input, " ") + 1
	ret.Tokens = make([]string, size)

	for index, element := range tokens {
		ret.Tokens[index] = element
	}

	ret.NameStr = ret.Tokens[0]
	ret.RawStr = input
	ret.Elements = size

	return ret
}
