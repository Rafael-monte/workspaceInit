package workspace

// Program Types
const navigatorType string = "Navegador"
const scriptType string = "Script"
const applicationType string = "Aplicacao"

var allProgramTypes = []string{navigatorType, scriptType, applicationType}

//Open browser window commands
var openWinBrowser = []string{"rundll32", "url.dll,FileProtocolHandler"}

const openLinuxBrowser string = "xdg-open"
const openMacBrowser string = "open"

//OS names
const win string = "windows"
const mac string = "darwin"
const linux string = "linux"

const configFileLocation string = "data\\config.json"
const emptyString string = ""
const directorySlash string = "/"
