package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/kr/pretty"
)

const root = ".git/refs/heads"

func main() {

	pretty.Logln("[INFO] initializing watcher...")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// listen
	var wg sync.WaitGroup
	pretty.Logln("[INFO] listening...")
	wg.Add(1)
	go func() {
		for {
			// handle event or error
			select {
			case event := <-watcher.Events:
				pretty.Logln("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					pretty.Logln("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				pretty.Logln("error:", err)
			}
		}
	}()

	// read file info from directory
	pretty.Logln("[INFO] reading directory...")
	fis, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	// iterate through the file infos
	pretty.Logln("[INFO] reading files...")
	for _, fi := range fis {

		// branch name is the file name
		branch := fi.Name()

		// add file to watchlist
		pretty.Logln("[INFO] adding", branch, "to watchlist...")
		err = watcher.Add(filepath.Join(root, branch))
		if err != nil {
			log.Fatal(err)
		}
	}

	wg.Wait()
}
