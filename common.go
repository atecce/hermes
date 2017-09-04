package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strings"
)

func execute(cmd *exec.Cmd) (string, error) {

	log.Println("[INFO] executing command", cmd.Path, cmd.Args)

	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout := outbuf.String()
	stderr := outbuf.String()

	if err != nil {

		// special case for provisioning
		if strings.Contains(stderr, "already exists") {
			log.Println("[INFO] looks like resource is already provisioned", errbuf.String())
			return stdout, resourceAlreadyProvisioned
		}

		// special case for cleaning
		if strings.Contains(err.Error(), "no such file or directory") {
			log.Println("[INFO] looks like tmp dir doesn't exist")
			return stdout, tmpDirDoesntExist
		}

		return stdout, errors.New(stderr)
	}

	log.Println("[INFO] stdout:\n", stdout)
	log.Println("[INFO] stderr:\n", stderr)

	return stdout, nil
}
