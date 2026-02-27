package fakecgo

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
	"testing"
)

func trimLastNewline(s string) string {
	if s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
		if s[len(s)-1] == '\r' {
			s = s[:len(s)-1]
		}
	}
	return s
}

func TestFailWithCgo(t *testing.T) {
	cmd := exec.CommandContext(t.Context(), "go", "build", "./")
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("command %v succeeded unexpectedly\n%s", cmd, out)
	}
	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) || !exitErr.Exited() {
		t.Fatalf("command %v failed: %v", cmd, err)
	}
	outString := trimLastNewline(string(out))
	if !regexp.MustCompile(`^\S+?: build constraints exclude all Go files in \S+?$`).MatchString(outString) {
		t.Fatalf("unexpected output: %s", outString)
	}
}
