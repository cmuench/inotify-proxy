package main

import (
	"flag"
	"github.com/cmuench/inotify-proxy/internal/watcher"
	"github.com/gookit/color"
	"strings"
)

var Version = "dev"

func main() {
	color.Style{color.FgWhite, color.FgDarkGray}.Printf("Version: %s\n", Version)

	sleepPtr := flag.Int("sleep", 2, "Cycle time in seconds. Defines time to sleep after each filesystem walk. Default 2s")
	profilePtr := flag.String("profile", "default", "Defines a special profile with extensions to look for. This speeds up the process. Available profiles are 'magento2-theme'")

	flag.Parse()

	includedDirectories := flag.Args()

	// If no argument is defined, the current directory is used
	if len(includedDirectories) == 0 {
		includedDirectories = append(includedDirectories, ".")
	}

	color.Style{color.FgCyan, color.OpBold}.Println("PROFILE: " + *profilePtr)
	color.Style{color.FgCyan, color.OpBold}.Println("DIRECTORIES: " + strings.Join(includedDirectories, ","))

	watcher.Watch(includedDirectories, *sleepPtr, *profilePtr)
}
