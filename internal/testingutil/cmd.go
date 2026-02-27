package testingutil

import (
	"os/exec"
	"testing"
)

func (t0 *T) Cmd(cmd *exec.Cmd) {
	t := (*testing.T)(t0)
	t.Helper()

	if cmd.Stdout == nil {
		cmd.Stdout = t.Output()
	}
	if cmd.Stderr == nil {
		cmd.Stderr = t.Output()
	}

	err := cmd.Run()
	if err != nil {
		t.Fatalf("command %v failed: %v", cmd, err)
	}
}
