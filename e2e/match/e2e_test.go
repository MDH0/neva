package test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	for i := 0; i < 10; i++ {
		cmd := exec.Command("neva", "run", "main")
		out, err := cmd.CombinedOutput()
		require.NoError(t, err)
		require.Equal(
			t,
			"one\ntwo\nthree\nfour\n",
			string(out),
		)
		require.Equal(t, 0, cmd.ProcessState.ExitCode())
	}
}
