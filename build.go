package main

import (
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

func build(name string) (string, error) {
	pretty.Logln("[INFO] building...")
	stdout, _ := execute(exec.Command("docker", "build", "-t", name, "-q", "."))
	ref := strings.Trim(stdout, "\n")
	return ref, nil
}
