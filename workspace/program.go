package workspace

import (
	"os/exec"
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

}

func (browserWindowProgram Program) openBrowserWindow() {
	const operationalSystem string = runtime.GOOS
	availableBrowsers := []browser{windowsBrowser{}, macBrowser{}, linuxBrowser{}}
	for _, browser := range availableBrowsers {
		browser.checkAndRunUrl(browserWindowProgram.Url, operationalSystem)
	}
}

func (applicationProgram Program) runApplication() {
	if applicationProgram.Type == scriptType {
		exec.Command(applicationProgram.Path, strings.Join(applicationProgram.OptionalArgs, " "))
		return
	}

	exec.Command(applicationProgram.Name)
}
