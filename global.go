package main

import (
	"flag"
	"os"
	"path/filepath"
)

// TODO break out into special implementation for launchd
var userAgentPath = filepath.Join(os.Getenv("HOME"), "Library", "LaunchAgents", "www.plist")
var (
	repoDir = flag.String("repo", os.Getenv("GIT_DIR"), "directory path for repository")
	tmpDir  = flag.String("tmp", os.TempDir(), "temporary directory path for build artefacts")

	// TODO make this configurable and get better default
	buildDir = filepath.Join(*tmpDir, "www")
)
