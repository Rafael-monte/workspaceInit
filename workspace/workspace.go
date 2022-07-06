package workspace

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//Conjunto de programas que serão rodados
type Workspace struct {
	//Nome do Workspace
	Name string `json:"workspaceName"`
	//Programas que irão ser rodados no workspace
	Programs []Program `json:"programs"`
}

func (workspace Workspace) RunPrograms() {
	for _, program := range workspace.Programs {
		fmt.Println("Executando", program.Name, "...")
		program.execute()
	}
}

var selectedWorkspace Workspace

//Inicia o programa, buscando os workspaces configurados e deixando o usuario escolher qual executar
func Create() Workspace {
	workspaces := getJSONFileInConfigurations()
	return selectWorkspaceFromOptions(workspaces)
}

func getJSONFileInConfigurations() (workspaces []Workspace) {

	filePath, filePathLocationError := filepath.Abs(configFileLocation)

	checkIfHasError(filePathLocationError, cannotFindConfigFilePath)

	jsonFile, readFileError := os.ReadFile(filePath)

	checkIfHasError(readFileError, cannotAccessConfigFile)

	conversionError := json.Unmarshal(jsonFile, &workspaces)

	checkIfHasError(conversionError, cannotConvertConfigFile)

	if len(workspaces) == 0 {
		log.Fatal(noWorkspacesSetted)
	}

	return
}

func selectWorkspaceFromOptions(workspaces []Workspace) (workspaceSelecionado Workspace) {
	if len(workspaces) > 1 {
		var opcao int
		fmt.Println("Selecione um workspace: ")
		for index, workspace := range workspaces {
			fmt.Println(index+1, workspace.Name)
		}
		fmt.Printf("opcao>: ")
		if _, err := fmt.Scanf("%d", &opcao); err != nil {
			log.Fatal(`Erro: Não foi possivel entender a entrada do usuário (Entrada esperada: Inteiro)`)
		}
		if opcao > len(workspaces)+1 || opcao <= 0 {
			log.Fatal(`Erro: A entrada esperada deve estar no intervalo entre 1 e`, len(workspaces)+1)
		}
		workspaceSelecionado = workspaces[opcao-1]
		selectedWorkspace = workspaceSelecionado
		return
	}
	workspaceSelecionado = workspaces[0]
	selectedWorkspace = workspaceSelecionado
	fmt.Println("Selecionando o unico workspace configurado:", workspaceSelecionado.Name)
	return
}

func GetSelectedWorkspace() Workspace {
	return selectedWorkspace
}
