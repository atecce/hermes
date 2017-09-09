package main

import (
	"os/exec"

	"github.com/kr/pretty"
)

func build(name string) (string, error) {
	pretty.Logln("[INFO] building...")
	return local{}.run(exec.Command("vagrant", "up"))
}
