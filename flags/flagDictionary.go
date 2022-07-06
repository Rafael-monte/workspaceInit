package flags

import (
	"fmt"
	"os"
	"workspaceInit/persistence"
	"workspaceInit/utils"
)

var postProgramFlags = map[string]func(){
	"--store":                 func() { persistence.SaveWorkspaceSelected("none") },
	"--store --override-data": func() { persistence.SaveWorkspaceSelected("--override-data") },
}

var flowChangerFlags = map[string]func(){
	"--info": func() {
		fmt.Println(`Workspace init é um inicializador de workspaces, feito para iniciar mais rápido os programas que você usa normalmente no dia a dia`)
	},

	"--teste": func() {
		fmt.Println("Isso aqui é só um teste mesmo!")
	},
}

// Checa se o usuario inseriu alguma flag que altera o fluxo principal, essas flags abrem novos que mudam completamente o comportamento da aplicacao
func CheckFlowChangerFlags() (functionStack []func()) {

	utils.StartLoadingResource("associação das flags ao contexto")
	usedFlags := []string{}
	os.Args = os.Args[1:]
	for _, flagValue := range os.Args {
		if alreadyUsedFlag, _ := utils.SliceContainsString(usedFlags, flagValue); !alreadyUsedFlag {
			if _, exists := flowChangerFlags[flagValue]; exists {
				usedFlags = append(usedFlags, flagValue)
				functionStack = append(functionStack, flowChangerFlags[flagValue])
			}
		}
	}
	utils.StopLoadingResource()
	return
}
