package env

import (
	"fmt"
	"os"
	"os/user"
)

type Vars struct {
	dirty bool
	count int
	env   map[string]string
}

func InitEnv() Vars {
	v := Vars{dirty: false, count: 0, env: nil}
	v.env = make(map[string]string)

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
	v.dirty = val
}

func (v *Vars) UpdateCount() {
	v.count = len(v.env)
}

func (v *Vars) GetCount() int {
	return v.count
}

func (v *Vars) CheckDirty() bool {
	return v.dirty
}

func (v *Vars) CheckExists(name string) bool {
	Var := v.env[name]
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
	v.env[name] = val
	v.UpdateCount()
}

func (v *Vars) SetEnvVar(name string, val string) {
	if v.CheckExists(name) == false {
		v.AddEnvVar(name, val)
		return
	} else {
		v.env[name] = val
		v.SetDirty(true)
	}
}

func (v *Vars) GetEnvVar(name string) string {
	if v.CheckExists(name) == true {
		return v.env[name]
	} else {
		return ""
	}
}

func (v *Vars) DeleteEnvVar(name string) bool {
	if v.CheckExists(name) == true {
		delete(v.env, name)
		v.UpdateCount()
		return true
	} else {
		return false
	}
}

func (v *Vars) Print() {
	fmt.Println("Total variables:", v.count, "\nDirty Environment:", v.dirty)
	for name, val := range v.env {
		fmt.Println(name, ":", val)
	}
}
