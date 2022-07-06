package main

import (
	"workspaceInit/flags"
	"workspaceInit/workspace"
)

var WorkspaceSelecionado workspace.Workspace

func main() {
	if alternativeFlow := flags.CheckFlowChangerFlags(); len(alternativeFlow) == 0 {
		// Main flow
		WorkspaceSelecionado = workspace.Create()
		WorkspaceSelecionado.RunPrograms()
	} else {
		for _, alternativeProcess := range alternativeFlow {
			alternativeProcess()
		}
	}
}
