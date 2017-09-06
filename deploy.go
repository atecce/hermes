package main

import (
	"os/exec"

	"github.com/kr/pretty"
)

func deploy(ref string) (string, error) {
	pretty.Logln("[INFO] deploying...")
	return execute(exec.Command("docker", "run", "-d", "-p", "8080:8080", ref))
}
