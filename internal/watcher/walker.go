package watcher

import (
	"github.com/cmuench/inotify-proxy/internal/profile/validator"
	"github.com/gookit/color"
	"os"
	"path/filepath"
	"time"
)

func Watch(includedDirectories []string, watchFrequenceSeconds int, profile string) {
	selectedProfile = profile

	for {
		for _, directoryToWalk := range includedDirectories {
			err := filepath.Walk(directoryToWalk, visit)

			if err != nil {
				panic(err)
			}
		}

		time.Sleep(time.Duration(watchFrequenceSeconds) * time.Second)
	}
}

func isFileChanged(path string, fileInfo os.FileInfo) bool {

	if !validator.IsPathValid(path, selectedProfile) {
		return false
	}

	nodeInfo, found := fileMap[path]

	currentModificationTime := fileInfo.ModTime()

	changed := false

	if !found {
		nodeInfo := NodeInfo{
			modificationUnixTime: currentModificationTime.Unix(),
		}
		fileMap[path] = nodeInfo

		color.Info.Println("Watching: " + path)
	} else {
		if nodeInfo.modificationUnixTime < currentModificationTime.Unix() {
			changed = true

			currentTime := time.Now()

			err := os.Chtimes(path, currentModificationTime, currentTime)

			if err != nil {
				panic("Error touching file" + path)
			}

			fileMap[path] = NodeInfo{
				modificationUnixTime: currentTime.Unix(),
			}
		}
	}

	return changed
}

func visit(path string, fileInfo os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return nil
	}

	fileChanged := isFileChanged(path, fileInfo)

	if fileChanged {
		color.Style{color.FgGreen, color.OpBold}.Printf("Changed: %s | %s\n", path, time.Now().Format("2006-01-02T15:04:05"))
	}

	return nil
}
