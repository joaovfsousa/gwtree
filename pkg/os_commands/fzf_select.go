package os_commands

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

const DEFAULT_SEARCH = ""

func FzfSelect(options []string, query *string) (string, error) {
	cmd := exec.Command("fzf")

	if query != nil {
		cmd = exec.Command("fzf", "--query", *query)
	}

	options_as_string := strings.Join(options, "\n")

	cmd.Stdin = bytes.NewReader([]byte(options_as_string))

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("Failed to execute command: %v", err)
	}

	lines := strings.Split(string(output), "\n")

	if len(lines) < 1 || lines[0] == "" {
		return "", errors.New("No options was selected")
	}

	return lines[0], nil
}
