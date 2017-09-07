package main

import (
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
	ref, err := getContainerRef("8080")
	_, err = execute(exec.Command("docker", "rm", "-f", ref))
	return err
}

func remoteDeploy(name string) error {
	pretty.Logln("[INFO] deploying remotely...")
	_, err := executeGce(exec.Command("sudo", "docker", "run", "-d", "-p", "80:8080", name))
	switch err {
	case portAlreadyAllocated:
		pretty.Logln("[INFO] looks like port is already allocated. cleaning...")
		if err := remoteClean(); err != nil {
			return err
		}
		return remoteDeploy(name)
	case nil:
		return nil
	}
	return err
}

func remoteClean() error {
	ref, err := getContainerRefGce("8080")
	_, err = executeGce(exec.Command("sudo", "docker", "rm", "-f", ref))
	return err
}

func getContainerRef(port string) (string, error) {
	stdout, err := execute(exec.Command("docker", "ps", "-f", "publish="+port, "-q"))
	if err != nil {
		return "", err
	}
	return strings.Trim(stdout, "\n"), nil
}

func getContainerRefGce(port string) (string, error) {
	stdout, err := executeGce(exec.Command("sudo", "docker", "ps", "-f", "publish="+port, "-q"))
	if err != nil {
		return "", err
	}
	return strings.Trim(stdout, "\n"), nil
}
