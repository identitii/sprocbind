package main_test

import (
	"testing"

	"github.com/identitii/sprocbind"
	"github.com/stretchr/testify/require"
)

func TestTypes(t *testing.T) {
	require.Equal(t, main.SQLTypeToGoType("NVARCHAR"), "string")
}
