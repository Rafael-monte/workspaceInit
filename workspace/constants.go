package workspace

// Program Types
const navigatorType string = "Navegador"
const appType string = "Aplicacao"
const scriptType string = "Script"

//Open browser window commands
var openWinBrowser = [2]string{"rundll32", "url.dll,FileProtocolHandler"}

const openLinuxBrowser string = "xdg-open"
const openMacBrowser string = "open"

//OS names
const win string = "windows"
const mac string = "darwin"
const linux string = "linux"

const configFileLocation string = "data\\config.json"
