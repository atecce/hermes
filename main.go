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

type playground struct {
	name    string
	current []byte // last observed head
	watcher *fsnotify.Watcher
}

func (p playground) watch() {

	defer p.watcher.Close()

	// add file to watchlist
	pretty.Logln("[INFO] adding", p.name, "to watchlist...")
	err := p.watcher.Add(p.name)
	if err != nil {
		log.Fatal(err)
	}

	pretty.Logln("[INFO] watching", p.name, "...")
	for {
		select {
		case event := <-p.watcher.Events:
			pretty.Logln("event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				pretty.Logln("modified file:", event.Name)
			}
		case err := <-p.watcher.Errors:
			pretty.Logln("error:", err)
		}
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	// read file info from directory
	pretty.Logln("[INFO] reading directory...")
	fis, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	// iterate through the file infos
	pretty.Logln("[INFO] reading files...")
	for _, fi := range fis {

		// get branch and head
		pretty.Logln("[INFO] getting branch and head...")
		branch := fi.Name()
		path := filepath.Join(root, branch)
		head, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		// create watcher
		pretty.Logln("[INFO] initializing watcher...")
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		p := playground{
			name:    path,
			current: head,
			watcher: watcher,
		}
		go p.watch()

	}

	wg.Wait()
}
