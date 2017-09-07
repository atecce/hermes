package main

import (
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

func build(name string) (string, error) {
	pretty.Logln("[INFO] building...")
	stdout, err := local{}.execute(exec.Command("docker", "build", "-t", name, "-q", "."))
	if err != nil {
		return stdout, err
	}
	ref := strings.Trim(stdout, "\n")
	return ref, nil
}
