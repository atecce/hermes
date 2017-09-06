package main

import (
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

func build() (string, error) {
	pretty.Logln("[INFO] building...")
	ref, err := execute(exec.Command("docker", "build", "-q", "."))
	return strings.Trim(ref, "\n"), err
}
