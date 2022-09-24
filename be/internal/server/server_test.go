package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitServer(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := require.New(t)

		s := InitServer()
		r.NotEmpty(s)
	})
}
