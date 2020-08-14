package profile

type Profile struct {
	fileExtensions []string
}

func (l *Profile) IsAllowedFileExtension(extension string) bool {
	for _, a := range l.fileExtensions {
		if a == extension {
			return true
		}
	}

	return false
}

var Magento2ThemeProfile = Profile{
	fileExtensions: []string{ ".css", ".js", ".less", ".sass" },
}

var Magento2 = Profile{
	fileExtensions: []string{ ".css", ".html", ".less", ".sass", ".js", ".php", ".phtml", ".xml" },
}

var VueStorefront = Profile{
	fileExtensions: []string{ ".css", ".js", ".sass", ".ts" },
}
