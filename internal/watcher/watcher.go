package watcher

import (
	"github.com/cmuench/inotify-proxy/internal/profile/validator"
	"github.com/cmuench/inotify-proxy/internal/util"
	"github.com/gookit/color"
	"github.com/karrick/godirwalk"
	"os"
	"time"
)

func visit(osPathname string, de *godirwalk.Dirent) error {
	// we only process files
	if de.IsDir() {
		return nil
	}

	if !validator.IsPathValid(osPathname, selectedProfile) {
		return godirwalk.SkipThis
	}

	fileChanged := isFileChanged(osPathname)
	if fileChanged {
		color.Style{color.FgGreen, color.OpBold}.Printf("Changed: %s | %s\n", osPathname, time.Now().Format("2006-01-02T15:04:05"))
	}

	return nil
}

func Watch(includedDirectories []string, watchFrequenceSeconds int, profile string) {
	selectedProfile = profile

	i := 0

	for {
		for _, directoryToWalk := range includedDirectories {
			err := godirwalk.Walk(directoryToWalk, &godirwalk.Options{
				Callback: visit,
				Unsorted: true,
			})

			if err != nil {
				panic(err)
			}
		}

		time.Sleep(time.Duration(watchFrequenceSeconds) * time.Second)

		if i%10 == 0 {
			garbageCollection()
			color.Info.Printf("Watching %d files ...\n", len(fileMap))
		}

		i++
	}
}

func isFileChanged(path string) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		color.Errorf("Cannot stat file %s\n", path)
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

			err := os.Chtimes(path, currentTime, currentModificationTime)

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

func garbageCollection() {
	for path, _ := range fileMap {
		if !util.FileExists(path) {
			delete(fileMap, path)
			color.Style{color.FgGray}.Printf("Deleted: %s\n", path)
		}
	}
}
