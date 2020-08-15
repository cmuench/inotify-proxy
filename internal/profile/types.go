package profile

import (
	"path/filepath"
	"strings"
)

type Profile struct {
	fileExtensions []string
}

func (l *Profile) IsAllowedFileExtension(path string) bool {

	// if profile contains only one extension definition with "*" then allow every extension.
	if len(l.fileExtensions) == 1 && l.fileExtensions[0] == "*" {
		return true
	}

	extension := filepath.Ext(path)

	for _, a := range l.fileExtensions {
		if a == extension {
			return true
		}
	}

	return false
}

func (l *Profile) IsAllowedDirectory(path string) bool {
	// Exclude some directories by default
	excludedDirectories := [...]string{
		"node_modules/",
		".idea/",
		".git/",
		".svn/",
	}

	for _, excludedDirectory := range excludedDirectories {
		if strings.Contains(path, excludedDirectory) {
			return false
		}
	}

	return true
}

var Default = Profile{
	fileExtensions: []string{"*"},
}

var Magento2Theme = Profile{
	fileExtensions: []string{".css", ".js", ".less", ".sass", ".ts"},
}

var Magento2 = Profile{
	fileExtensions: []string{".css", ".html", ".less", ".sass", ".js", ".php", ".phtml", ".ts", ".xml"},
}

var VueStorefront = Profile{
	fileExtensions: []string{".css", ".js", ".sass", ".ts"},
}
