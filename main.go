package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

const gitDir = "/home/git"

var (
	tmpDir  = filepath.Join(gitDir, "tmp/www")
	repoDir = filepath.Join(gitDir, "www.git")
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
	return executeCommand("rm", "-rf", tmpDir)
}

func clone() error {
	log.Println("[INFO] cloning...")
	return executeCommand("git", "clone", repoDir, tmpDir)
}

func build() error {
	log.Println("[INFO] building...")
	return executeCommand("jekyll", "build", "-s", tmpDir, "-d", filepath.Join(tmpDir, "_site"))
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

func deploy() error {
	log.Println("[INFO] deploying...")
	return executeCommand("gcloud", "compute", "copy-files", "/home/git/tmp/www/_site", "atec@atec:/home/atec", "--zone", "us-east1-b")
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
		log.Println("[ERROR] failed to provision (maybe instance is already up?)")
		log.Println(err)
	}
	err = deploy()
	if err != nil {
		log.Println("[FATAL] failed to deploy")
		log.Fatal(err)
	}
}

func executeCommand(argv ...string) error {

	cmd := exec.Command(argv[0], argv[1:]...)

	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	if err := cmd.Run(); err != nil {

		if strings.Contains(errbuf.String(), "already exists") {
			return resourceAlreadyProvisioned
		}

		output := outbuf.String() + errbuf.String()
		return errors.New(output)
	}
	return nil
}
