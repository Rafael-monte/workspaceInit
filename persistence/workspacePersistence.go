package persistence

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"workspaceInit/workspace"
)

var currentWorkspace workspace.Workspace

func SaveWorkspaceSelected(options string) {
	currentWorkspace = workspace.GetSelectedWorkspace()

	fileInByteArray, compressionError := json.MarshalIndent(currentWorkspace, "", " ")

	if compressionError != nil {
		log.Fatal(`Erro: Não foi possivel compactar os dados em JSON para salvá-los`)
	}

	var persistenceError error = nil
	if options == overrideData {
		persistenceError = ioutil.WriteFile(persistencePlace, fileInByteArray, fs.FileMode(allPermissionsCode))
	}

	if persistenceError != nil {
		log.Fatal(`Erro: Não foi possivel persistir workspace`, currentWorkspace.Name, `(override-data: `, options == overrideData, `)`)
	}

}
