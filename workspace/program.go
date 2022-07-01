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
	switch program.Type {
	case navigatorType:
		program.openBrowserWindow()
	case scriptType:
		program.runScript()
	case applicationType:
		program.runApplication()
	default:
		log.Fatal(`Erro: O tipo`, program.Type, `não existe, verifique a gramática e tente novamente.
		Os tipos permitidos são: 
		`, strings.Join(allProgramTypes, "\n"))
	}
}

func (browserWindowProgram Program) openBrowserWindow() {
	const operationalSystem string = runtime.GOOS
	availableBrowsers := []browser{windowsBrowser{}, macBrowser{}, linuxBrowser{}}
	for _, browser := range availableBrowsers {
		browser.checkAndRunUrl(browserWindowProgram.Url, operationalSystem)
	}
}

func (scriptProgram Program) runScript() {
	absFilePath := getAbsolutePath(scriptProgram.Path)
	if scriptProgram.Type == scriptType {
		scriptCmd := exec.Command(absFilePath, strings.Join(scriptProgram.OptionalArgs, " ")).Start()
		if scriptCmd.Error() != emptyString {
			fmt.Println(scriptCmd.Error())
		}
		return
	}
}

func (applicationProgram Program) runApplication() {
	absoluteFilePath := getAbsolutePath(applicationProgram.Path)
	var appCmdError error = nil
	if strings.Contains(applicationProgram.Path, directorySlash) {
		appCmdError = exec.Command(absoluteFilePath).Start()
	} else {
		appCmdError = exec.Command(applicationProgram.Path).Start()
	}

	if appCmdError.Error() != emptyString {
		fmt.Println(appCmdError.Error())
	}
}

func getAbsolutePath(path string) string {
	absFilePath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(`Erro: Não foi possivel determinar o caminho da aplicação/Script (`, path, `)`)
	}
	return absFilePath
}
