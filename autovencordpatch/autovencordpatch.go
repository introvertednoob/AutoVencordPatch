package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	discordJSON   = "/Applications/Discord.app/Contents/Resources/build_info.json"
	vencordApp    = "/Applications/VencordInstaller.app"
	checkInterval = 1 * time.Second
)

func runInstaller() {
	cmd := exec.Command("open", vencordApp)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to run installer:", err)
	}
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Failed to create watcher:", err)
		return
	}
	defer watcher.Close()

	dir := filepath.Dir(discordJSON)
	err = watcher.Add(dir)
	if err != nil {
		fmt.Println("Failed to add watcher:", err)
		return
	}

	fmt.Println("Watching Discord version...")

	for {
		select {
		case event := <-watcher.Events:
			if filepath.Clean(event.Name) == discordJSON && event.Op&fsnotify.Create == fsnotify.Create {
				time.Sleep(1.0 * time.Second)
				runInstaller()
			}
		case err := <-watcher.Errors:
			fmt.Println("Watcher error:", err)
			time.Sleep(checkInterval)
		}
	}
}