package watcher

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"testing"
	"time"
)

type sysDir struct {
	name  string
	files []string
}

var sysdir = func() *sysDir {
	switch runtime.GOOS {
	case "darwin":
		switch runtime.GOARCH {
		case "arm64":
			wd, err := syscall.Getwd()
			if err != nil {
				wd = err.Error()
			}
			sd := &sysDir{
				filepath.Join(wd, "..", ".."),
				[]string{
					"ResourceRules.plist",
					"Info.plist",
				},
			}
			found := true
			for _, f := range sd.files {
				path := filepath.Join(sd.name, f)
				if _, err := os.Stat(path); err != nil {
					found = false
					break
				}
			}
			if found {
				return sd
			}
			// In a self-hosted iOS build the above files might
			// not exist. Look for system files instead below.
		}
	case "windows":
		return &sysDir{
			os.Getenv("SystemRoot") + "\\system32\\drivers\\etc",
			[]string{
				"networks",
				"protocol",
				"services",
			},
		}
	case "plan9":
		return &sysDir{
			"/lib/ndb",
			[]string{
				"common",
				"local",
			},
		}
	}
	return &sysDir{
		"/etc",
		[]string{
			"passwd",
			"group",
			"hosts",
		},
	}
}()

func TestIfFileChangedNotModified(t *testing.T) {
	var sfdir = sysdir.name
	var sfname = sysdir.files[0]

	// First call registers file
	assert.False(t, isFileChanged(filepath.Join(sfdir, sfname)))

	// Second call checks if file is changed
	assert.False(t, isFileChanged(filepath.Join(sfdir, sfname)))
}

func TestIfFileChangedModified(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "inotify-watch-test-")
	if err != nil {
		t.Fatal("Cannot create temporary file")
	}

	// First call registers temp file
	assert.False(t, isFileChanged(tmpFile.Name()))

	currentTime := time.Now()

	// simulate change
	chtimesError := os.Chtimes(tmpFile.Name(), currentTime, currentTime)

	if chtimesError != nil {
		t.Fatal("Cannot chtimes of temporary file")
	}

	assert.False(t, isFileChanged(tmpFile.Name()))
}
