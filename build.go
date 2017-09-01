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
	return executeCommand("jekyll", "build", "-s", buildDir, "-d", filepath.Join(buildDir, "_site"))
}

func buildSvc() error {
	log.Println("[INFO] building web server...")
	return executeCommand("go", "build", filepath.Join(buildDir, "main.go"))
}
