package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

type deployer interface {
	deploy(name string) error
}

func localDeploy(name string) error {
	pretty.Logln("[INFO] deploying locally...")
	_, err := execute(exec.Command("docker", "run", "-d", "-p", "8080:8080", name))
	switch err {
	case portAlreadyAllocated:
		pretty.Logln("[INFO] looks like port is already allocated. cleaning...")
		if err := localClean(); err != nil {
			return err
		}
		return localDeploy(name)
	case nil:
		return nil
	}
	return err
}

func localClean() error {
	stdout, err := execute(exec.Command("docker", "ps", "-f", "publish=8080", "-q"))
	if err != nil {
		return err
	}
	ref := strings.Trim(stdout, "\n")
	_, err = execute(exec.Command("docker", "rm", "-f", ref))
	return err
}

func remoteDeploy(name string) error {
	pretty.Logln("[INFO] deploying remotely...")
	cmd := exec.Command("gcloud", "compute", "ssh", "atec", "--zone", "us-east1-b", "--command", fmt.Sprintf("'sudo docker run -d -p 80:8080 %s'", name))
	_, err := execute(cmd)
	return err
}
