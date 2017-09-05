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

// TODO break out into special implementation for launchd
var userAgentPath = filepath.Join(os.Getenv("HOME"), "Library", "LaunchAgents", "www.plist")
var (
	repoDir = flag.String("repo", os.Getenv("GIT_DIR"), "directory path for repository")
	tmpDir  = flag.String("tmp", os.TempDir(), "temporary directory path for build artefacts")

	// TODO make this configurable and get better default
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

func deploy() error {
	pretty.Logln("[INFO] deploying...")
	cmd := exec.Command("launchctl", "unload", userAgentPath)
	if _, err := execute(cmd); err != nil {
		pretty.Logln("[ERROR] failed to unload launchd job, assuming it hasn't started...")
	}
	cmd = exec.Command("launchctl", "load", userAgentPath)
	if _, err := execute(cmd); err != nil {
		return err
	}
	return nil
}

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
	err = deploy()
	if err != nil {
		pretty.Logln("[FATAL] failed to deploy")
		log.Fatal(err)
	}
}
