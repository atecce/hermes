package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
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
	return executeCommand("rm", "-rf", "/home/git/tmp/www")
}

func clone() error {
	log.Println("[INFO] cloning...")
	return executeCommand("git", "clone", "/home/git/www.git", "/home/git/tmp/www")
}

func build() error {
	log.Println("[INFO] building...")
	return executeCommand("jekyll", "build", "-s", "/home/git/tmp/www", "-d", "/home/git/tmp/www/_site")
}

func test() error {
	log.Println("[INFO] testing...")
	return nil
}

var resourceAlreadyProvisioned = errors.New("resource already provisioned")

func provision() error {
	log.Println("[INFO] provisioning...")
	return executeCommand("gcloud", "compute", "instances", "create", "atec", "--zone", "us-east1-b", "--format", "json")
}

func deploy() error {
	log.Println("[INFO] deploying...")
	return executeCommand("gcloud", "compute", "copy-files", "/home/git/tmp/www/_site", "atec@atec:/home/atec", "--zone", "us-east1-b", "--format", "json")
}

func main() {
	err := clean()
	if err != nil {
		log.Println("[FATAL] failed to clean")
		log.Fatal("[FATAL] ", err)
	}
	err = clone()
	if err != nil {
		log.Println("[FATAL] failed to clone")
		log.Fatal("[FATAL] ", err)
	}
	err = build()
	if err != nil {
		log.Println("[FATAL] failed to build")
		log.Fatal("[FATAL] ", err)
	}
	err = provision()
	if err != nil {
		log.Println("[ERROR] failed to provision (maybe instance is already up?)")
		log.Println("[ERROR] ", err)
	}
	err = deploy()
	if err != nil {
		log.Println("[FATAL] failed to deploy")
		log.Fatal("[FATAL] ", err)
	}
}

func executeCommand(argv ...string) error {

	cmd := exec.Command(argv[0], argv[1:]...)

	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	if err := cmd.Run(); err != nil {

		output := outbuf.String() + "\n" + errbuf.String()
		return errors.New(output)
	}
	return nil
}
