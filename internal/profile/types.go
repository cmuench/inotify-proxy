package profile

type Profile struct {
	Extensions []string
}

var Default = Profile{
	Extensions: []string{"*"},
}

var LESS = Profile{
	Extensions: []string{".less"},
}

var Magento2Theme = Profile{
	Extensions: []string{".css", ".js", ".less", ".sass", ".ts"},
}

var Magento2 = Profile{
	Extensions: []string{".css", ".html", ".less", ".sass", ".js", ".php", ".phtml", ".ts", ".xml"},
}

var SASS = Profile{
	Extensions: []string{".sass", ".scss"},
}

var VueStorefront = Profile{
	Extensions: []string{".css", ".js", ".sass", ".ts"},
}

var Javascript = Profile{
	Extensions: []string{".js", ".ts"},
}
