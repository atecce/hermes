package main

import (
	"log"
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
	return executeCommand("jekyll", "build", "-s", tmpDir, "-d", filepath.Join(tmpDir, "_site"))
}

func buildSvc() error {
	log.Println("[INFO] building web server...")
	return executeCommand("GOOS=linux", "GOARCH=386", "go", "build", filepath.Join(tmpDir, "main.go"))
}
