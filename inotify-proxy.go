package main

import (
	"flag"
	"github.com/cmuench/inotify-proxy/internal/config"
	"github.com/cmuench/inotify-proxy/internal/util"
	"github.com/cmuench/inotify-proxy/internal/watcher"
	"github.com/gookit/color"
	"os"
	"strings"
)

// Version defines the version of the application. This variable will be overridden by build system
var Version = "dev"

func main() {
	color.Style{color.FgWhite, color.FgDarkGray}.Printf("Version: %s\n", Version)

	sleepPtr := flag.Int("sleep", 2, "Cycle time in seconds. Defines time to sleep after each filesystem walk. Default 2s")
	profilePtr := flag.String("profile", "default", "Defines a special profile with extensions to look for. This speeds up the process. Available profiles are 'magento2-theme'")
	noConfig := flag.Bool("no-config", false, "Do not load config.")

	flag.Parse()

	includedDirectories := flag.Args()
	c := config.Config{}

	if !*noConfig {
		if util.FileExists("inotify-proxy.yaml") {

			r, err := os.Open("inotify-proxy.yaml")

			if err != nil {
				color.Errorf("cannot read file: %v\n", err)
				os.Exit(1)
			}

			defer r.Close()

			c, err = config.Read(r)

			if err != nil {
				color.Errorf("cannot read config: %v\n", err)
			}

			if c.OldGlobalProfile != nil {
				color.Errorf("You are using the old configuration format. Please use the new configuration version.\n")
				color.Print("\nPlease refer: https://github.com/cmuench/inotify-proxy/blob/master/README.md#config\n")
				os.Exit(1)
			}
		}
	}

	if len(includedDirectories) > 0 {
		for _, includedDirectory := range includedDirectories {
			c.Entries = append(c.Entries, config.WatchEntry{
				Directory:  includedDirectory,
				Extensions: nil,
				Profile:    profilePtr,
			})
		}
	}

	// If no argument is defined, the current directory is used
	if len(c.Entries) == 0 {
		c.AddEntry(config.WatchEntry{
			Directory:  ".",
			Extensions: nil,
			Profile:    profilePtr,
		})
	}

	color.Style{color.FgMagenta, color.OpBold}.Println("Watching ...")
	color.Style{color.FgWhite}.Println(strings.Repeat("-", 80))

	for _, e := range c.Entries {
		color.Style{color.FgCyan, color.OpBold}.Printf("Directory: %s\n", e.Directory)
		if *e.Profile != "" {
			color.Style{color.FgCyan, color.OpBold}.Printf("Profile: %s\n", *e.Profile)
		}
		if len(e.Extensions) > 0 {
			color.Style{color.FgCyan, color.OpBold}.Printf("Extensions: %s\n", e.Extensions)
		}

		color.Style{color.FgWhite}.Println(strings.Repeat("-", 80))
	}

	watcher.Watch(c, *sleepPtr)
}
