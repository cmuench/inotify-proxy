package profile

// Profile is the base type of any profile
type Profile struct {
	Extensions []string
}

// Default is the default profile which allows all extensions
var Default = Profile{
	Extensions: []string{"*"},
}

// LESS allows only .less extension
var LESS = Profile{
	Extensions: []string{".less"},
}

// Magento2Theme allows only extensions of a Magento 2 Theme
var Magento2Theme = Profile{
	Extensions: []string{".css", ".js", ".less", ".sass", ".ts"},
}

// Magento2 allows only extensions required for development of a Magento 2 Module / Extension
var Magento2 = Profile{
	Extensions: []string{".css", ".html", ".less", ".sass", ".js", ".php", ".phtml", ".ts", ".xml"},
}

// SASS allow only extensions related to the SASS tool
var SASS = Profile{
	Extensions: []string{".sass", ".scss"},
}

// VueStorefront allow only extensions related to vue-storefront projects
var VueStorefront = Profile{
	Extensions: []string{".css", ".js", ".sass", ".ts"},
}

// Javascript allow only .js (Javascript) and .ts (Typescript) extension
var Javascript = Profile{
	Extensions: []string{".js", ".ts"},
}
