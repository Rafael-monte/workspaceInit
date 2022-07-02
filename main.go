package main

import "workspaceInit/workspace"

var WorkspaceSelecionado workspace.Workspace

func main() {
	WorkspaceSelecionado = workspace.Create()
	WorkspaceSelecionado.RunPrograms()
}
