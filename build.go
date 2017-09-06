package main

import (
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

func build() (string, error) {
	pretty.Logln("[INFO] building...")
	stdout, _ := execute(exec.Command("docker", "build", "-q", "."))
	ref := strings.Trim(stdout, "\n")
	return ref, nil
}
