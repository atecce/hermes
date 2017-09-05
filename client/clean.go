package main

import (
	"errors"
	"os/exec"

	"github.com/kr/pretty"
)

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
