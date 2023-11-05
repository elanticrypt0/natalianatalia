package helpers

import (
	"os/exec"
	"runtime"
)

func OpenInBrowser(url string) {

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", "explorer", url)
	} else {
		cmd = exec.Command("open", url)
	}

	cmd.Run()
}
