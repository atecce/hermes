package main

import (
	"log"
	"os/exec"
	"path/filepath"
)

func build() error {
	log.Println("[INFO] building...")
	if err := buildHtml(); err != nil {
		return err
	}
	if err := buildSvc(); err != nil {
		return err
	}
	return nil
}

func buildHtml() error {
	log.Println("[INFO] building html docs...")
	cmd := exec.Command("/usr/local/bin/jekyll", "build", "-s", buildDir, "-d", filepath.Join(buildDir, "_site"))
	return executeCommand(cmd)
}

func buildSvc() error {
	log.Println("[INFO] building web server...")
	cmd := exec.Command("/usr/bin/go", "build", filepath.Join(buildDir, "main.go"))
	return executeCommand(cmd)
}
