package main

import "workspaceInit/workspace"

func main() {
	workspaceSelecionado := workspace.Create()
	workspaceSelecionado.RunPrograms()
}
