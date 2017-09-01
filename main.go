package main

import (
	"errors"
	"log"
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

func clean() error {
	log.Println("[INFO] cleaning...")
	return executeCommand("rm", "-rf", buildDir)
}

func clone() error {
	log.Println("[INFO] cloning...")
	return executeCommand("git", "clone", repoDir, buildDir)
}

func test() error {
	log.Println("[INFO] testing...")
	return nil
}

var resourceAlreadyProvisioned = errors.New("resource already provisioned")

func provision() error {
	log.Println("[INFO] provisioning...")
	err := executeCommand("gcloud", "compute", "instances", "create", "atec", "--zone", "us-east1-b")
	switch err {
	case resourceAlreadyProvisioned:
		log.Println("[INFO] resource already provisioned. skipping")
		return nil
	default:
		return err
	}
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
	err = deploy()
	if err != nil {
		log.Println("[FATAL] failed to deploy")
		log.Fatal(err)
	}
}
