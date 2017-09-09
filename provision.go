package main

import (
	"os/exec"

	"github.com/kr/pretty"
)

func provision() error {
	pretty.Logln("[INFO] provisioning...")
	cmd := exec.Command("vagrant", "provision")
	_, err := local{}.run(cmd)
	return err
}
