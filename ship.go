package main

import (
	"os/exec"

	"github.com/kr/pretty"
)

func ship(name string) error {
	pretty.Logln("[INFO] shipping...")
	_, err := local{}.run(exec.Command("docker", "push", name))
	if err != nil {
		return err
	}
	return nil
}
