package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func executeCommand(argv ...string) error {

	cmd := exec.Command(argv[0], argv[1:]...)
	cmd.Env = []string{"GOOS=linux", "GOARCH=386"}
	//	cmd.Dir = buildDir

	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	if err := cmd.Run(); err != nil {

		if strings.Contains(errbuf.String(), "already exists") {
			return resourceAlreadyProvisioned
		}

		output := outbuf.String() + errbuf.String()
		return errors.New(output)
	}
	return nil
}
