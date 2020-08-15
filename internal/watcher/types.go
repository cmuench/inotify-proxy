package watcher

type NodeInfo struct {
	modificationUnixTime int64
}

var fileMap = make(map[string]NodeInfo)

var selectedProfile = ""
