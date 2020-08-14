package main

import (
	"flag"
	"github.com/cmuench/inotify-proxy/internal/profile"
	"github.com/gookit/color"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type NodeInfo struct {
	modificationUnixTime int64
}

var fileMap = make(map[string]NodeInfo)
var selectedProfile = ""

func shouldSkipFile(path string) bool {

	fileExtension := filepath.Ext(path)

	// Exclude some directories by default

	if strings.Contains(path, "node_modules/") {
		return true
	}

	if strings.Contains(path, ".idea/") {
		return true
	}

	// Check profiles

	if selectedProfile == "" {
		return false
	}

	if selectedProfile == "magento2-theme" {
		if profile.Magento2ThemeProfile.IsAllowedFileExtension(fileExtension) {
			return false
		}
	}

	if selectedProfile == "magento2" {
		if profile.Magento2.IsAllowedFileExtension(fileExtension) {
			return false
		}
	}

	if selectedProfile == "vue-storefront" {
		if profile.VueStorefront.IsAllowedFileExtension(fileExtension) {
			return false
		}
	}

	return true
}

func isFileChanged(path string, fileInfo os.FileInfo) bool {

	if shouldSkipFile(path) {
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


func main() {
	sleepPtr := flag.Int("sleep", 2, "Cycle time in seconds. Defines time to sleep after each filesystem walk. Default 2s")
	profilePtr := flag.String("profile", "", "Defines a special profile with extensions to look for. This speeds up the process. Available profiles are 'magento2-theme'")

	flag.Parse()

	includedDirectories := flag.Args()

	selectedProfile = *profilePtr

	// If no argument is defined, the current directory is used
	if len(includedDirectories) == 0 {
		includedDirectories = append(includedDirectories, ".")
	}

	color.Style{color.FgCyan, color.OpBold}.Println("PROFILE: " + selectedProfile)
	color.Style{color.FgCyan, color.OpBold}.Println("DIRECTORIES: " + strings.Join(includedDirectories, ","))

	for {

		for _, directoryToWalk := range includedDirectories {
			err := filepath.Walk(directoryToWalk, visit)

			if err != nil {
				panic(err)
			}
		}

		time.Sleep(time.Duration(*sleepPtr) * time.Second)
	}
}
