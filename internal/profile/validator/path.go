package validator

import (
	"github.com/cmuench/inotify-proxy/internal/config"
	"github.com/cmuench/inotify-proxy/internal/profile"
	"path/filepath"
	"strings"
)

func IsPathValid(path string, entryConfig config.WatchEntry) bool {

	if !isAllowedDirectory(path) {
		return false
	}

	if len(entryConfig.Extensions) > 0 && !isAllowedFileExtension(path, entryConfig.Extensions) {
		 return false
	}

	if entryConfig.Profile == nil {
		return true
	}

	var selectedProfile profile.Profile

	switch *entryConfig.Profile {
	case "less":
		selectedProfile = profile.LESS
	case "magento2":
		selectedProfile = profile.Magento2
	case "magento2-theme":
		selectedProfile = profile.Magento2Theme
	case "sass":
		selectedProfile = profile.SASS
	case "vue-storefront":
		selectedProfile = profile.VueStorefront
	case "javascript":
		selectedProfile = profile.Javascript
	default:
		selectedProfile = profile.Default
	}

	return isAllowedFileExtension(path, selectedProfile.Extensions)
}

func isAllowedDirectory(path string) bool {
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

func isAllowedFileExtension(path string, fileExtensions []string) bool {

	// if profile contains only one extension definition with "*" then allow every extension.
	if len(fileExtensions) == 1 && fileExtensions[0] == "*" {
		return true
	}

	extension := filepath.Ext(path)

	for _, a := range fileExtensions {
		if a == extension {
			return true
		}
	}

	return false
}
