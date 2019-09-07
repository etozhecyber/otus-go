package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	require.Equal(t, unpack("a4bc2d5e"), "aaaabccddddde")
	require.Equal(t, unpack("abcd"), "abcd")
	require.Equal(t, unpack("45"), "")
	require.Equal(t, unpack("a45"), "aaaa")
}
