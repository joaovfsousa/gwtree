package os_commands

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func ExecOsCmd(name string, arg ...string) ([]string, error) {
	cmd := exec.Command(name, arg...)

	output, err := cmd.Output()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to execute command: %v", err))
	}

	lines := strings.Split(string(output), "\n")

	return lines, nil
}
