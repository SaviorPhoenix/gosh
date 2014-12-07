package shell

import (
	"github.com/SaviorPhoenix/gosh/env"
	"os"
	"os/signal"
)

type Shell struct {
	SigChan     chan os.Signal
	Environment env.Vars
}

var Sh Shell

func (sh *Shell) InitSignalHandler() {
	sh.SigChan = make(chan os.Signal, 32)
	signal.Notify(sh.SigChan)
}

func (sh *Shell) InitShell() {
	sh.InitSignalHandler()
	sh.Environment = env.InitEnv()
}

func (sh *Shell) GetEnv() env.Vars {
	return sh.Environment
}
