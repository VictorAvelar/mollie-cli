package utils

import (
	"os/exec"
	"runtime"

	"github.com/sirupsen/logrus"
)

// Browse opens the desired target in your native default browser.
func Browse(target string) {
	switch runtime.GOOS {
	case "linux":
		logrus.Fatal(exec.Command("xdg-open", target).Start())
	case "windows":
		logrus.Fatal(exec.Command("rundll32", "url.dll,FileProtocolHandler", target).Start())
	case "darwin":
		logrus.Fatal(exec.Command("open", target).Start())
	default:
		// for the BSDs
		logrus.Fatal(exec.Command("xdg-open", target).Start())
	}
}
