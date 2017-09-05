package main

import (
	"os/exec"

	"github.com/kr/pretty"
)

func clone() error {
	pretty.Logln("[INFO] cloning...")
	cmd := exec.Command("git", "clone", ".", buildDir)
	if _, err := execute(cmd); err != nil {
		return err
	}
	return nil
}
