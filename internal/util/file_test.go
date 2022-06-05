package util

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	// Test for current file_test.go
	_, file, _, _ := runtime.Caller(1)
	assert.True(t, FileExists(file))
	assert.False(t, FileExists(file+".not-existing"))
}
