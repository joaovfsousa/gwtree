package os_commands

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecOsCmd(name string, arg ...string) ([]string, error) {
	cmd := exec.Command(name, arg...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Failed to execute command: %v: %v: %v", cmd.String(), err.Error(), string(output))
	}

	lines := strings.Split(string(output), "\n")

	return lines, nil
}
