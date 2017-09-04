package main

import (
	"os/exec"
	"path/filepath"

	"github.com/kr/pretty"
)

func configure() error {
	pretty.Logln("[INFO] configuring...")
	cmd := exec.Command("gcloud", "compute", "copy-files", filepath.Join(buildDir, "_site"), "atec@atec:/home/atec", "--zone", "us-east1-b")
	if _, err := execute(cmd); err != nil {
		return err
	}
	return nil
}
