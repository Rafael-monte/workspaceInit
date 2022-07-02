package flags

import (
	"workspaceInit/persistence"
)

var postProgramFlags = map[string]func(){
	"--store":                 func() { persistence.SaveWorkspaceSelected("none") },
	"--store --override-data": func() { persistence.SaveWorkspaceSelected("--override-data") },
}

var flowChangerFlags = map[string]func(){
	"--info": func() {},
}

func CheckFlags() {

}
