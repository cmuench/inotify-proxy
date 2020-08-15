package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultProfile(t *testing.T) {
	selectedProfile := ""
	assert.True(t, IsPathValid("README.md", selectedProfile))
}

func TestMagentoProfile(t *testing.T) {
	selectedProfile := "magento2"

	assert.False(t, IsPathValid("README.md", selectedProfile))
	assert.True(t, IsPathValid("foo.js", selectedProfile))
	assert.True(t, IsPathValid("foo.ts", selectedProfile))
	assert.True(t, IsPathValid("foo.php", selectedProfile))
	assert.True(t, IsPathValid("foo.phtml", selectedProfile))
	assert.True(t, IsPathValid("foo.html", selectedProfile))
}
