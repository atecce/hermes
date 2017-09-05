package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/kr/pretty"
)

var (
	repoDir = flag.String("repo", os.Getenv("GIT_DIR"), "directory path for repository")
	tmpDir  = flag.String("tmp", os.TempDir(), "temporary directory path for build artefacts")

	// TODO remove special case
	buildDir = filepath.Join(*tmpDir, "www")
)

type temp interface {
	clone() error
	build() error
	test() error
	provision() error
	configure() error
	deploy() error
	monitor() error
}

var tmpDirDoesntExist = errors.New("tmp dir doesn't exist")

func clean() error {
	pretty.Logln("[INFO] cleaning...")
	cmd := exec.Command("rm", "-rf", buildDir)
	_, err := execute(cmd)
	switch err {
	case tmpDirDoesntExist:
		pretty.Logln("[INFO] build dir doesn't exist. skipping")
		return nil
	default:
		return err
	}
}

func clone() error {
	pretty.Logln("[INFO] cloning...")
	cmd := exec.Command("git", "clone", ".", buildDir)
	if _, err := execute(cmd); err != nil {
		return err
	}
	return nil
}

func test() error {
	pretty.Logln("[INFO] testing...")
	return nil
}

var userAgent = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>www</string>
	<key>Program</key>
	<string>/tmp/www/main</string>
</dict>
</plist>`)

func main() {
	err := clean()
	if err != nil {
		pretty.Logln("[FATAL] failed to clean")
		log.Fatal(err)
	}
	err = clone()
	if err != nil {
		pretty.Logln("[FATAL] failed to clone")
		log.Fatal(err)
	}
	err = build()
	if err != nil {
		pretty.Logln("[FATAL] failed to build")
		log.Fatal(err)
	}
	err = configure()
	if err != nil {
		pretty.Logln("[FATAL] failed to configure")
		log.Fatal(err)
	}
}
