package utils

import (
	"os/exec"
	"runtime"
)

// Browse opens the desired target in your native default browser.
func Browse(target string) {
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", target).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", target).Start()
	case "darwin":
		exec.Command("open", target).Start()
	default:
		// for the BSDs
		exec.Command("xdg-open", target).Start()
	}
}
