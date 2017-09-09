package main

import (
	"os/exec"

	"github.com/kr/pretty"
)

func configure(name string) error {
	pretty.Logln("[INFO] configuring...")
	_, err := remote{}.run(exec.Command("sudo", "docker", "pull", name))
	if err != nil {
		return err
	}
	return nil
}
