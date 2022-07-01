package workspace

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//Programa que vai ser iniciado na rotina de abertura do workspace
type Program struct {
	//Nome do programa (Ex: Google Chrome)
	Name string `json:"name"`
	//Caminho da aplicacao no sistema
	Path string `json:"path"`
	//Tipo de aplicacao: Pode ser "Navegador", "Aplicacao" ou "Script"
	Type string `json:"type"`
	//Parametros opcionais do programa, geralmente usado quando é um script
	OptionalArgs []string `json:"optArgs"`
	//Url do programa, utilizado quando o programa é do tipo navegador
	Url string `json:"URL"`
}

func (program Program) execute() {
	if program.Type == navigatorType {
		program.openBrowserWindow()
		return
	}
	program.runApplication()

}

func (browserWindowProgram Program) openBrowserWindow() {
	const operationalSystem string = runtime.GOOS
	availableBrowsers := []browser{windowsBrowser{}, macBrowser{}, linuxBrowser{}}
	for _, browser := range availableBrowsers {
		browser.checkAndRunUrl(browserWindowProgram.Url, operationalSystem)
	}
}

func (applicationProgram Program) runApplication() {
	absFilePath, err := filepath.Abs(applicationProgram.Path)
	if err != nil {
		log.Fatal(`Erro: Não foi possivel determinar o caminho da aplicação/Script (`, applicationProgram, `)`)
	}

	if applicationProgram.Type == scriptType {
		scriptCmd := exec.Command(absFilePath, strings.Join(applicationProgram.OptionalArgs, " ")).Start()
		if scriptCmd.Error() != emptyString {
			fmt.Println(scriptCmd.Error())
		}
		return
	}

	appCmd := exec.Command(absFilePath).Start()
	if appCmd.Error() != emptyString {
		fmt.Println(appCmd.Error())
	}
}
