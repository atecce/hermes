package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"git.atec.pub/sh"
	"github.com/fsnotify/fsnotify"
	"github.com/kr/pretty"
)

// TODO
const (
	root     = ".git/refs/heads"
	testPath = "/Users/atec/go/src/atec.pub/www"
)

type playground struct {
	name    string
	current []byte // last observed head
	watcher *fsnotify.Watcher
}

func New(branch string) playground {
	pretty.Logf("[INFO] initializing playground for %s...", branch)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	head, err := ioutil.ReadFile(filepath.Join(root, branch))
	if err != nil {
		log.Fatal(err)
	}
	return playground{
		name:    filepath.Join(root, branch),
		current: head,
		watcher: watcher,
	}
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
			go sh.Run(exec.Command("bundle", "exec", "jekyll", "serve", "-s", testPath+"/.hermes/master", "-d", testPath+"/.hermes/master/_site"))

		case err := <-p.watcher.Errors:
			pretty.Logln("[ERROR]", err)
		}
	}
}

func readDir(root string) []os.FileInfo {
	pretty.Logln("[INFO] reading directory...")
	fis, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	return fis
}

func main() {

	// iterate through the file infos
	pretty.Logln("[INFO] reading files...")
	fis := readDir(root)
	for _, fi := range fis {
		p := New(fi.Name())
		go p.watch()

	}

	// block forever
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
