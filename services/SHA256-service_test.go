package services_test

import (
	"go-auth/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSHA256Service(t *testing.T) {
	t.Parallel()

	string, err := services.SHA256Encoder("my-value")
	require.NoError(t, err)
	require.NotNil(t, string)

	string2, err2 := services.SHA256Encoder("")
	require.Empty(t, string2)
	require.Error(t, err2)
}
