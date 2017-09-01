package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strings"
)

func execute(cmd *exec.Cmd) error {

	log.Println("[INFO] executing command", cmd.Path, cmd.Args)

	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	output := outbuf.String() + errbuf.String()
	log.Println("[INFO] combined output:", output)

	if err != nil {

		log.Println("[ERROR]", err)

		// special case for provisioning
		if strings.Contains(errbuf.String(), "already exists") {
			log.Println("[INFO] looks like resource is already provisioned", errbuf.String())
			return resourceAlreadyProvisioned
		}

		// special case for cleaning
		if strings.Contains(err.Error(), "no such file or directory") {
			log.Println("[INFO] looks like tmp dir doesn't exist")
			return tmpDirDoesntExist
		}

		return errors.New(output)
	}

	return nil
}
