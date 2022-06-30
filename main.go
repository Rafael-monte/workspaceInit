package main

import "workspaceInit/workspace"

func main() {
	workspaceSelecionado := workspace.Run()
	workspaceSelecionado.RunPrograms()
}
