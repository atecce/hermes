package main

import (
	"log"
	"os/exec"
	"path/filepath"
)

func configure() error {
	log.Println("[INFO] configuring...")
	cmd := exec.Command("gcloud", "compute", "copy-files", filepath.Join(buildDir, "_site"), "atec@atec:/home/atec", "--zone", "us-east1-b")
	return execute(cmd)
}
