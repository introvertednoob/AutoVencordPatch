package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	cacheFile     = "cached.txt"
	discordJSON   = "/Applications/Discord.app/Contents/Resources/build_info.json"
	vencordApp    = "/Applications/VencordInstaller.app"
	checkInterval = 1 * time.Second // fallback interval
)

func readCachedVersion() string {
	if _, err := os.Stat(cacheFile); os.IsNotExist(err) {
		err := ioutil.WriteFile(cacheFile, []byte("0.0.0"), 0644)
		if err != nil {
			fmt.Println("Failed to create cache file:", err)
		}
		return "0.0.0"
	}

	data, err := ioutil.ReadFile(cacheFile)
	if err != nil {
		fmt.Println("Failed to read cache file:", err)
		return "0.0.0"
	}

	return string(data)
}

func writeCachedVersion(version string) {
	err := ioutil.WriteFile(cacheFile, []byte(version), 0644)
	if err != nil {
		fmt.Println("Failed to write cache file:", err)
	}
}

func readDiscordVersion() (string, error) {
	data, err := ioutil.ReadFile(discordJSON)
	if err != nil {
		return "", err
	}

	var parsed struct {
		Version string `json:"version"`
	}

	if err := json.Unmarshal(data, &parsed); err != nil {
		return "", err
	}

	return parsed.Version, nil
}

func runInstaller() {
	cmd := exec.Command("open", vencordApp)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to run installer:", err)
	}
}

func main() {
	cachedVersion := readCachedVersion()

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
				currentVersion, err := readDiscordVersion()
				if err != nil {
					fmt.Println("Failed to read Discord version:", err)
					continue
				}
				if currentVersion != cachedVersion {
					cachedVersion = currentVersion
					writeCachedVersion(cachedVersion)
					time.Sleep(1.0 * time.Second)
					runInstaller()
				}
			}
		case err := <-watcher.Errors:
			fmt.Println("Watcher error:", err)
			time.Sleep(checkInterval)
		}
	}
}