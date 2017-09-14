package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"atec.pub/sh"
	"github.com/fsnotify/fsnotify"
	"github.com/kr/pretty"
)

const root = ".git/refs/heads"

type playground struct {
	name    string
	current []byte // last observed head
	watcher *fsnotify.Watcher
}

// TODO
const testPath = "/Users/atec/go/src/atec.pub/www"

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

			pretty.Logln("[INFO] rcvd event:", event)
			pretty.Logln("[INFO] modified file:", event.Name)
			next, err := ioutil.ReadFile(p.name)
			if err != nil {
				log.Fatal(err)
			}
			pretty.Logln("[INFO] old head:", strings.Trim(string(p.current), "\n"))
			pretty.Logln("[INFO] new head:", strings.Trim(string(next), "\n"))
			p.current = next

			// TODO
			sh.Run(exec.Command("rm", "-rf", testPath+"/.hermes/master"))
			sh.Run(exec.Command("git", "clone", testPath, testPath+"/.hermes/master"))
			sh.Run(exec.Command("bundle", "exec", "jekyll", "build", "-s", testPath+"/.hermes/master", "-d", testPath+"/.hermes/master/_site"))

		case err := <-p.watcher.Errors:
			pretty.Logln("[ERROR]", err)
		}
	}
}

func main() {

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

	// block forever
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
