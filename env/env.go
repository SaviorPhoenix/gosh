package env

import (
	"fmt"
	"os"
	"os/user"
)

type Vars struct {
	Dirty bool
	Count int
	Env   map[string]string
}

func InitEnv() Vars {
	v := Vars{Dirty: false, Count: 0, Env: nil}
	v.Env = make(map[string]string)

	usr, _ := user.Current()

	v.AddEnvVar("prompt", "$ ")
	v.AddEnvVar("home", usr.HomeDir)
	v.AddEnvVar("editor", os.Getenv("EDITOR"))
	v.AddEnvVar("term", os.Getenv("TERM"))
	v.AddEnvVar("pager", os.Getenv("PAGER"))
	v.AddEnvVar("user", os.Getenv("USER"))
	v.UpdateCount()
	return v
}

func (v *Vars) SetDirty(val bool) {
	v.Dirty = val
}

func (v *Vars) UpdateCount() {
	v.Count = len(v.Env)
}

func (v *Vars) CheckExists(name string) bool {
	Var := v.Env[name]
	if Var != "" {
		return true
	} else {
		return false
	}
}

func (v *Vars) AddEnvVar(name string, val string) {
	if v.CheckExists(name) == true {
		v.SetEnvVar(name, val)
		return
	}
	v.Env[name] = val
	v.UpdateCount()
}

func (v *Vars) SetEnvVar(name string, val string) {
	if v.CheckExists(name) == false {
		v.AddEnvVar(name, val)
		return
	} else {
		v.Env[name] = val
		v.SetDirty(true)
	}
}

func (v *Vars) GetEnvVar(name string) string {
	if v.CheckExists(name) == true {
		return v.Env[name]
	} else {
		return ""
	}
}

func (v *Vars) DeleteEnvVar(name string) bool {
	if v.CheckExists(name) == true {
		delete(v.Env, name)
		v.UpdateCount()
		return true
	} else {
		return false
	}
}

func (v *Vars) Print() {
	fmt.Println("Total variables:", v.Count, "\nDirty Environment:", v.Dirty)
	for name, val := range v.Env {
		fmt.Println(name, ":", val)
	}
}
