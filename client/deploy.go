package main

import (
	"os/exec"

	"github.com/kr/pretty"
)

func deploy() error {
	pretty.Logln("[INFO] deploying...")
	cmd := exec.Command("launchctl", "unload", userAgentPath)
	if _, err := execute(cmd); err != nil {
		pretty.Logln("[ERROR] failed to unload launchd job, assuming it hasn't started...")
	}
	cmd = exec.Command("launchctl", "load", userAgentPath)
	if _, err := execute(cmd); err != nil {
		return err
	}
	return nil
}
