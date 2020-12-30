package validator

import (
	"github.com/cmuench/inotify-proxy/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtensionAllowList(t *testing.T) {
	we := config.WatchEntry{
		Extensions: []string{".js", ".less"},
	}

	// Node modules are excluded
	assert.False(t, IsPathValid("foo/node_modules/test.js", we))
	assert.False(t, IsPathValid("foo/node_modules/test.less", we))

	// Node modules are excluded
	assert.True(t, IsPathValid("example/test.js", we))
	assert.True(t, IsPathValid("example/test.less", we))

	assert.False(t, IsPathValid("README.md", we))
	assert.False(t, IsPathValid(".git/config", we))
	assert.False(t, IsPathValid("foo/.git/config", we))

}

func TestMagentoProfile(t *testing.T) {
	p := "magento2"
	we := config.WatchEntry{
		Extensions: nil,
		Profile:    &p,
	}

	assert.False(t, IsPathValid("foo/node_modules/test.js", we))
	assert.False(t, IsPathValid(".git/config", we))
	assert.False(t, IsPathValid("foo/.git/config", we))

	assert.False(t, IsPathValid("README.md", we))
	assert.False(t, IsPathValid("test.zip", we))
	assert.True(t, IsPathValid("foo.js", we))
	assert.True(t, IsPathValid("foo.ts", we))
	assert.True(t, IsPathValid("foo.php", we))
	assert.True(t, IsPathValid("foo.phtml", we))
	assert.True(t, IsPathValid("foo.html", we))
}

func TestDefaultProfile(t *testing.T) {
	p := "default"
	we := config.WatchEntry{
		Extensions: nil,
		Profile:    &p,
	}

	assert.False(t, IsPathValid("foo/node_modules/test.js", we))
	assert.False(t, IsPathValid(".git/config", we))
	assert.False(t, IsPathValid("foo/.git/config", we))

	assert.True(t, IsPathValid("README.md", we))
	assert.True(t, IsPathValid("foo.js", we))
	assert.True(t, IsPathValid("foo.ts", we))
	assert.True(t, IsPathValid("foo.php", we))
	assert.True(t, IsPathValid("foo.phtml", we))
	assert.True(t, IsPathValid("foo.html", we))
	assert.True(t, IsPathValid("foo.xyz", we))
}

func TestSASSProfile(t *testing.T) {
	p := "sass"
	we := config.WatchEntry{
		Extensions: nil,
		Profile:    &p,
	}

	assert.False(t, IsPathValid("foo/node_modules/test.js", we))
	assert.False(t, IsPathValid(".git/config", we))
	assert.False(t, IsPathValid("foo/.git/config", we))

	assert.True(t, IsPathValid("foo.sass", we))
	assert.True(t, IsPathValid("foo.scss", we))
}
