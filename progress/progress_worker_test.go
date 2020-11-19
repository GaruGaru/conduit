package progress

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {

	worker := NewAsciiProgressWorker(20)

	for i := 1; i <= 20; i++ {
		worker.SetCurrent(int64(i))
	}

	worker.Finish()

	require.Equal(t, int64(20), worker.GetCurrent())
}
