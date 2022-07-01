package workspace

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

//Inicia o programa, buscando os workspaces configurados e deixando o usuario escolher qual executar
func Create() Workspace {
	workspaces := getJSONFileInConfigurations()
	return selectWorkspaceFromOptions(workspaces)
}

func getJSONFileInConfigurations() (workspaces []Workspace) {
	// absPath, err := filepath.Abs("../data/createdWorkspaces.json")
	// if err != nil {
	// 	log.Fatal(`Não foi possivel converter o caminho do arquivo de configurações para absoluto`)
	// }
	jsonFile, err := os.ReadFile(configFileLocation)
	if err != nil {
		log.Fatal(`Erro: Não foi possivel acessar o arquivo de configuracoes dos workspaces.
			Por favor, certifique-se que:
			- O arquivo de configurações existe (Localizado em: "data/createdWorkspaces.json")
			- O arquivo está formatado corretamente
		`)
	}

	err = json.Unmarshal(jsonFile, &workspaces)
	if err != nil {
		log.Fatal(`Erro: Não foi possivel converter o arquivo de de configurações dos workspaces
			Possivelmente o arquivo foi formatado incorretamente
		`)
	}

	if len(workspaces) == 0 {
		log.Fatal(`Erro: Não existe nenhum workspace configurado para execução. Crie um e tente novamente`)
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
		return
	}
	workspaceSelecionado = workspaces[0]
	fmt.Println("Selecionando o unico workspace configurado:", workspaceSelecionado.Name)
	return
}
