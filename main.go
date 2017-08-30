package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
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

func clone() error {
	log.Println("[INFO] cloning...")
	err := exec.Command("rm", "-rf", "/home/git/tmp/www").Run()
	if err != nil {
		return err
	}
	return exec.Command("git", "clone", "/home/git/www.git", "/home/git/tmp/www").Run()
}

func build() error {
	log.Println("[INFO] building...")
	return exec.Command("jekyll", "build", "-s", "/home/git/tmp/www", "-d", "/home/git/tmp/www/_site").Run()
}

func test() error {
	log.Println("[INFO] testing...")
	return nil
}

func provision() error {
	log.Println("[INFO] provisioning...")
	cmd := exec.Command("gcloud", "compute", "instances", "create", "atec", "--zone", "us-east1-b")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		// TODO check nested err?
		stderr, _ := ioutil.ReadAll(os.Stderr)
		return errors.New(string(stderr))
	}
	return nil
}

func deploy() error {
	log.Println("[INFO] deploying...")
	err := exec.Command("gcloud", "compute", "scp", "notebook", "atec.pub:/home/atec/").Run()
	if err != nil {
		return err
	}
	return exec.Command("gcloud", "compute", "scp", "_site/*", "atec.pub:/var/www/").Run()
}

func main() {
	err := clone()
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
		log.Println("[FATAL] failed to provision")
		log.Fatal("[FATAL] ", err)
	}
}
