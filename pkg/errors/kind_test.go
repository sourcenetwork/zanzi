package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ErrorKind_SameKind_ReturnsIsTrue(t *testing.T) {
	require.ErrorIs(t, BadInput, BadInput)
}

func Test_ErrorKind_DifferentKind_ReturnsFalse(t *testing.T) {
	require.False(t, errors.Is(BadInput, Internal))
}

func Test_ErrorKind_DifferentType_ReturnsFalse(t *testing.T) {
	err := errors.New("test")
	require.False(t, errors.Is(err, Internal))
}
