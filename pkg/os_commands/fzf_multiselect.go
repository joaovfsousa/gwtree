package os_commands

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func FzfMultiselect(options []string) ([]string, error) {
	cmd := exec.Command("fzf", "-m")

	options_as_string := strings.Join(options, "\n")

	cmd.Stdin = bytes.NewReader([]byte(options_as_string))

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed to execute command: %v", err)
	}

	lines := strings.Split(string(output), "\n")

	if len(lines) < 1 || lines[0] == "" {
		return nil, errors.New("No options was selected")
	}

	return lines[0 : len(lines)-1], nil
}
