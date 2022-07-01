package workspace

import (
	"os/exec"
)

type browser interface {
	openUrl(url string)
	getOS() string
	checkAndRunUrl(os, url string)
}

type windowsBrowser struct{}

type macBrowser struct{}

type linuxBrowser struct{}

func (browser windowsBrowser) openUrl(url string) {
	exec.Command(openWinBrowser[0], openWinBrowser[1], url).Start()
}

func (browser macBrowser) openUrl(url string) {
	exec.Command(openMacBrowser, url).Start()
}

func (browser linuxBrowser) openUrl(url string) {
	exec.Command(openLinuxBrowser, url).Start()
}

func (browser windowsBrowser) getOS() string {
	return win
}

func (browser macBrowser) getOS() string {
	return mac
}

func (browser linuxBrowser) getOS() string {
	return linux
}

func (browser windowsBrowser) checkAndRunUrl(url, os string) {
	if os == browser.getOS() {
		browser.openUrl(url)
	}
}

func (browser macBrowser) checkAndRunUrl(url, os string) {
	if os == browser.getOS() {
		browser.openUrl(url)
	}
}
func (browser linuxBrowser) checkAndRunUrl(url, os string) {
	if os == browser.getOS() {
		browser.openUrl(url)
	}
}
