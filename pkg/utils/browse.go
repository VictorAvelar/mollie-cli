package utils

import (
	"os/exec"
	"runtime"

	"github.com/sirupsen/logrus"
)

// Browse opens the desired target in your native default browser.
func Browse(target string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", target).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", target).Start()
	case "darwin":
		err = exec.Command("open", target).Start()
	default:
		// for the BSDs
		err = exec.Command("xdg-open", target).Start()
	}

	if err != nil {
		logrus.Fatal(err)
	}
}
