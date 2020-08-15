package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMagentoProfile(t *testing.T) {
	selectedProfile := "magento2"

	assert.False(t, IsPathValid("foo/node_modules/test.js", selectedProfile))
	assert.False(t, IsPathValid(".git/config", selectedProfile))
	assert.False(t, IsPathValid("foo/.git/config", selectedProfile))

	assert.False(t, IsPathValid("README.md", selectedProfile))
	assert.False(t, IsPathValid("test.zip", selectedProfile))
	assert.True(t, IsPathValid("foo.js", selectedProfile))
	assert.True(t, IsPathValid("foo.ts", selectedProfile))
	assert.True(t, IsPathValid("foo.php", selectedProfile))
	assert.True(t, IsPathValid("foo.phtml", selectedProfile))
	assert.True(t, IsPathValid("foo.html", selectedProfile))
}

func TestDefaultProfile(t *testing.T) {
	selectedProfile := "default"

	assert.False(t, IsPathValid("foo/node_modules/test.js", selectedProfile))
	assert.False(t, IsPathValid(".git/config", selectedProfile))
	assert.False(t, IsPathValid("foo/.git/config", selectedProfile))

	assert.True(t, IsPathValid("README.md", selectedProfile))
	assert.True(t, IsPathValid("foo.js", selectedProfile))
	assert.True(t, IsPathValid("foo.ts", selectedProfile))
	assert.True(t, IsPathValid("foo.php", selectedProfile))
	assert.True(t, IsPathValid("foo.phtml", selectedProfile))
	assert.True(t, IsPathValid("foo.html", selectedProfile))
	assert.True(t, IsPathValid("foo.xyz", selectedProfile))
}
