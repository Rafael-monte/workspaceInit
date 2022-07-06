package utils

import (
	"fmt"
	"time"
)

var loading = true
var eventMarker chan struct{} = make(chan struct{}) //marcador de evento

const loadingDefaultText = "Carregando"
const numberOfFrames time.Duration = 24

//Utilizado para criar uma string de carregamento
func StartLoadingResource(baseText string) {
	go showLoadingText(baseText)
}

func StopLoadingResource() {
	loading = false
	<-eventMarker
}

func showLoadingText(baseText string) {
	for loading {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c %s `%s`", r, loadingDefaultText, baseText)
			time.Sleep(time.Second / numberOfFrames)
		}
	}
	fmt.Printf("\r \r")
	eventMarker <- struct{}{}
}
