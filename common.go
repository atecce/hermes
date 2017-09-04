package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

func execute(cmd *exec.Cmd) (string, error) {

	pretty.Logln("[INFO] executing command", cmd.Path, cmd.Args)

	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout := outbuf.String()
	stderr := errbuf.String()

	pretty.Logln("[INFO] stdout:\n\n", stdout)
	pretty.Logln("[INFO] stderr:\n\n", stderr)

	if err != nil {

		// special case for provisioning
		if strings.Contains(stderr, "already exists") {
			pretty.Logln("[INFO] looks like resource is already provisioned\n\n", stderr)
			return stdout, resourceAlreadyProvisioned
		}

		return stdout, errors.New(stderr)
	}

	return stdout, nil
}
