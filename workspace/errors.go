package workspace

import "log"

const cannotFindConfigFilePath string = `Erro: Não foi possivel encontrar o caminho do arquivo de configurações`

const cannotAccessConfigFile string = `Erro: Não foi possivel acessar o arquivo de configuracoes dos workspaces.
Por favor, certifique-se que:
- O arquivo de configurações existe (Localizado em: "data/createdWorkspaces.json")
- O arquivo está formatado corretamente
`

const cannotConvertConfigFile string = `Erro: Não foi possivel converter o arquivo de de configurações dos workspaces
Possivelmente o arquivo foi formatado incorretamente
`

const noWorkspacesSetted string = `Erro: Não existe nenhum workspace configurado para execução. Crie um e tente novamente`

func checkIfHasError(possibleError error, errorMessage string) {
	if possibleError != nil {
		log.Fatal(errorMessage)
	}
}
