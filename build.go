package main

import (
	"os/exec"
	"path/filepath"

	"github.com/kr/pretty"
)

var siteDir = filepath.Join(buildDir, "_site")

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
	cmd := exec.Command("jekyll", "build", "-s", buildDir, "-d", siteDir)
	if _, err := execute(cmd); err != nil {
		return err
	}
	return nil
}

func buildSvc() error {
	pretty.Logln("[INFO] building web server...")
	cmd := exec.Command("go", "build", filepath.Join(buildDir, "main.go"))
	cmd.Dir = siteDir
	if _, err := execute(cmd); err != nil {
		return err
	}
	return nil
}
