package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ErrorIs_WithCause_ReturnsTrueIfCauseIsTheSame(t *testing.T) {
	cause := errors.New("cause")
	err := NewWithCause("test", Internal, cause)
	require.ErrorIs(t, err, cause)
}

func Test_ErrorIs_WithCause_ReturnsFalseIfCauseIsNotTheSame(t *testing.T) {
	notCause := errors.New("anything")
	err := NewWithCause("test", Internal, errors.New("cause"))
	require.False(t, errors.Is(err, notCause))
}

func Test_ErrorIs_TrueIfKindIsSame(t *testing.T) {
	require.ErrorIs(t, New("test", Internal), Internal)
}
func Test_ErrorIs_FalseWithDifferentKind(t *testing.T) {
	require.False(t, errors.Is(New("test", Internal), BadInput))
}

func Test_Wrap_WrappedErrorIsPreserved(t *testing.T) {
	base := New("something", Internal)
	err := Wrap("more things", base)
	require.ErrorIs(t, err, base)
}

func Test_Wrap_WrappedErrorKind_IsOfKind(t *testing.T) {
	err := Wrap("more things", Internal)
	require.ErrorIs(t, err, Internal)
}

func Test_Wrap_WrapNonZanziError_KindIsUnknown(t *testing.T) {
	base := errors.New("something")
	err := Wrap("more things", base)
	require.ErrorIs(t, err, Unknown)
}
