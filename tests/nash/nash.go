package nash

import (
	"os"
	"testing"

	"github.com/NeowayLabs/nash"
)

func Setup(t *testing.T) *nash.Shell {
	shell, err := nash.New()

	if err != nil {
		t.Fatal(err)
	}

	shell.SetStdout(os.Stdout)
	shell.SetStderr(os.Stderr)
	nashPath := os.Getenv("HOME") + "/.nash"
	os.MkdirAll(nashPath, 0655)
	shell.SetDotDir(nashPath)

	return shell
}
