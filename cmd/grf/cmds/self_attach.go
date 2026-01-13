package cmds

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

// AttachSelf create new grf process to attach to caller's pid
func AttachSelf(filename string) error {
	// check grf first
	binName := "grf"
	if runtime.GOOS == "windows" {
		binName = "grf.exe"
	}

	grfPath, err := exec.LookPath(binName)
	if err != nil {
		return fmt.Errorf("grf not found in PATH: %v", err)
	}

	// execute grf
	cmd := exec.Command(grfPath, "attach", strconv.Itoa(os.Getpid()), "-o", filename)
	return cmd.Run()
}
