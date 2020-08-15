package validator

import (
	"github.com/cmuench/inotify-proxy/internal/profile"
)

func IsPathValid(path string, profileName string) bool {

	var selectedProfile profile.Profile

	switch profileName {
	case "magento2":
		selectedProfile = profile.Magento2
	case "magento2-theme":
		selectedProfile = profile.Magento2Theme
	case "vue-storefront":
		selectedProfile = profile.VueStorefront
	default:
		selectedProfile = profile.Default
	}

	return selectedProfile.IsAllowedDirectory(path) && selectedProfile.IsAllowedFileExtension(path)
}
