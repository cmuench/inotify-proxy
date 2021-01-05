package watcher

// NodeInfo contains the required information for a file to watch
type NodeInfo struct {
	modificationUnixTime int64
}

var fileMap = make(map[string]NodeInfo)
