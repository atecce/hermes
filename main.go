package main

import (
	"errors"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/kr/pretty"
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
	cmd := exec.Command("git", "clone", repoDir, buildDir)
	if _, err := execute(cmd); err != nil {
		return err
	}
	return nil
}

func test() error {
	pretty.Logln("[INFO] testing...")
	return nil
}

var resourceAlreadyProvisioned = errors.New("resource already provisioned")

func provision() error {
	pretty.Logln("[INFO] provisioning...")
	cmd := exec.Command("gcloud", "compute", "instances", "create", "atec", "--zone", "us-east1-b", "--tags", "http-server")
	_, err := execute(cmd)
	switch err {
	case resourceAlreadyProvisioned:
		pretty.Logln("[INFO] resource already provisioned. skipping")
		return nil
	default:
		return err
	}
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
	err = provision()
	if err != nil {
		pretty.Logln("[ERROR] failed to provision")
		pretty.Logln(err)
	}
	err = configure()
	if err != nil {
		pretty.Logln("[FATAL] failed to configure")
		log.Fatal(err)
	}
}
