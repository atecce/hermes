package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kr/pretty"
)

func configure() error {

	pretty.Logln("[INFO] configuring...")

	userAgentsDir := filepath.Join(os.Getenv("HOME"), "Library", "LaunchAgents", "www.plist")
	if err := ioutil.WriteFile(userAgentsDir, userAgent, 0644); err != nil {
		pretty.Logln("[ERROR]", err)
		return errors.New("failed to write launchd User Agent")
	}
	return nil
}
