package main

import (
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

func deploy(ref string) error {
	pretty.Logln("[INFO] deploying...")
	_, err := execute(exec.Command("docker", "run", "-d", "-p", "8080:8080", ref))
	switch err {
	case portAlreadyAllocated:
		pretty.Logln("[INFO] looks like port is already allocated. cleaning...")
		if err := clean(); err != nil {
			return err
		}
		return deploy(ref)
	case nil:
		return nil
	}
	return err
}

func clean() error {
	stdout, err := execute(exec.Command("docker", "ps", "-f", "publish=8080", "-q"))
	if err != nil {
		return err
	}
	ref := strings.Trim(stdout, "\n")
	_, err = execute(exec.Command("docker", "rm", "-f", ref))
	return err
}
