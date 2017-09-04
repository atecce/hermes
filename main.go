package main

import (
	"errors"
	"log"
	"os/exec"
	"path/filepath"
)

const gitDir = "/home/git"

var (
	repoDir  = filepath.Join(gitDir, "www.git")
	tmpDir   = filepath.Join(gitDir, "tmp")
	buildDir = filepath.Join(tmpDir, "www")
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
	log.Println("[INFO] cleaning...")
	cmd := exec.Command("rm", "-rf", buildDir)
	err := execute(cmd)
	switch err {
	case tmpDirDoesntExist:
		log.Println("[INFO] build dir doesn't exist. skipping")
		return nil
	default:
		return err
	}
}

func clone() error {
	log.Println("[INFO] cloning...")
	cmd := exec.Command("git", "clone", repoDir, buildDir)
	return execute(cmd)
}

func test() error {
	log.Println("[INFO] testing...")
	return nil
}

func main() {
	err := clean()
	if err != nil {
		log.Println("[FATAL] failed to clean")
		log.Fatal(err)
	}
	err = clone()
	if err != nil {
		log.Println("[FATAL] failed to clone")
		log.Fatal(err)
	}
	err = build()
	if err != nil {
		log.Println("[FATAL] failed to build")
		log.Fatal(err)
	}
	err = provision()
	if err != nil {
		log.Println("[ERROR] failed to provision")
		log.Println(err)
	}
	err = configure()
	if err != nil {
		log.Println("[FATAL] failed to configure")
		log.Fatal(err)
	}
}
