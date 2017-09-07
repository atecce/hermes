package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

func execute(cmd *exec.Cmd) (string, error) {

	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout := outbuf.String()
	stderr := errbuf.String()

	if *verbose {
		pretty.Logln("[INFO] executing command", cmd.Path, cmd.Args, "...")
		pretty.Logln("[INFO] stdout:\n\n", stdout)
		pretty.Logln("[INFO] stderr:\n\n", stderr)
	}

	if err != nil {
		return stdout, checkErr(stderr)
	}
	return stdout, nil
}

var (
	portAlreadyAllocated       = errors.New("port is already allocated") // docker
	resourceAlreadyProvisioned = errors.New("resource already exists")   // gce
	tmpDirDoesntExist          = errors.New("tmp dir doesn't exist")     // unix fs
)

func checkErr(stderr string) error {
	if strings.Contains(stderr, "port is already allocated") {
		return portAlreadyAllocated
	}
	if strings.Contains(stderr, "already exists") {
		return resourceAlreadyProvisioned
	}
	return errors.New(stderr)
}
