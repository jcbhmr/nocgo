package testingutil

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func (t0 *T) Eval(code string) string {
	t := (*testing.T)(t0)
	t.Helper()

	err := os.MkdirAll(".test/eval", 0o755)
	if err != nil {
		t.Fatalf("failed to create eval directory %q: %v", ".test/eval", err)
	}
	t.Cleanup(func() {
		_ = os.RemoveAll(".test/eval")
	})

	fName := func() string {
		f, err := os.CreateTemp(".test/eval", "eval-*.go")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		t.Cleanup(func() {
			_ = os.Remove(f.Name())
		})
		defer f.Close()

		_, err = f.WriteString(code)
		if err != nil {
			t.Fatalf("failed to write code to temp file %q: %v", f.Name(), err)
		}

		return f.Name()
	}()

	cmd := exec.Command("go", "fmt", fName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to format code: %v\n%s", err, out)
	}

	o := fName + ".out"
	if runtime.GOOS == "windows" {
		o += ".exe"
	}
	cmd = exec.Command("go", "build", "-o", o, fName)
	out, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to build code: %v\n%s", err, out)
	}
	t.Cleanup(func() {
		_ = os.Remove(o)
	})

	cmd = exec.Command(o)
	out, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run code: %v\n%s", err, out)
	}

	outString := string(out)
	if strings.HasSuffix(outString, "\r\n") {
		outString = outString[:len(outString)-2]
	} else if strings.HasSuffix(outString, "\n") {
		outString = outString[:len(outString)-1]
	}
	return outString
}
