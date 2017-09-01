package main

import (
	"log"
	"os/exec"
)

func deploy() error {
	log.Println("[INFO] deploying...")
	if err := deploySite(); err != nil {
		return err
	}
	if err := deploySvc(); err != nil {
		return err
	}
	return nil
}

func deploySite() error {
	log.Println("[INFO] deploying site...")
	cmd := exec.Command("gcloud", "compute", "copy-files", "/home/git/tmp/www/_site", "atec@atec:/home/atec", "--zone", "us-east1-b")
	return executeCommand(cmd)
}

func deploySvc() error {
	log.Println("[INFO] deploying service...")
	return nil
}
