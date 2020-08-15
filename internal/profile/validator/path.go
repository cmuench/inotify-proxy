package validator

import (
	"github.com/cmuench/inotify-proxy/internal/profile"
	"path/filepath"
	"strings"
)

func IsPathValid(path string, selectedProfile string) bool {

	fileExtension := filepath.Ext(path)

	// Exclude some directories by default

	if strings.Contains(path, "node_modules/") {
		return false
	}

	if strings.Contains(path, ".idea/") {
		return false
	}

	switch selectedProfile {
	case "magento2":
		return profile.Magento2.IsAllowedFileExtension(fileExtension)
	case "magento2-theme":
		return profile.Magento2Theme.IsAllowedFileExtension(fileExtension)
	case "vue-storefront":
		return profile.VueStorefront.IsAllowedFileExtension(fileExtension)
	default:
		return true
	}
}
