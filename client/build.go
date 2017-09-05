package main

import (
	"os/exec"
	"path/filepath"

	"github.com/kr/pretty"
)

func build() error {
	pretty.Logln("[INFO] building...")
	if err := buildHtml(); err != nil {
		return err
	}
	if err := buildSvc(); err != nil {
		return err
	}
	return nil
}

func buildHtml() error {
	pretty.Logln("[INFO] building html docs...")
	cmd := exec.Command("jekyll", "build", "-s", ".", "-d", buildDir)
	cmd.Dir = filepath.Join(*repoDir, "../") // TODO this assumes post-commit hook
	if _, err := execute(cmd); err != nil {
		return err
	}
	return nil
}

func buildSvc() error {
	pretty.Logln("[INFO] building web server...")
	cmd := exec.Command("go", "build", "-o", filepath.Join(buildDir, "main"), "main.go")
	cmd.Dir = filepath.Join(*repoDir, "../") // TODO this assumes post-commit hook
	if _, err := execute(cmd); err != nil {
		return err
	}
	return nil
}
