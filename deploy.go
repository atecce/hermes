package main

import (
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

type deployer interface {
	deploy(name string) error
}

func (_ local) deploy(name string) error {
	pretty.Logln("[INFO] deploying locally...")
	_, err := local{}.run(exec.Command("docker", "run", "-d", "-p", "8080:8080", name))
	switch err {
	case portAlreadyAllocated:
		pretty.Logln("[INFO] looks like port is already allocated. cleaning...")
		if err := (local{}.clean()); err != nil {
			return err
		}
		return local{}.deploy(name)
	case nil:
		return nil
	}
	return err
}

type cleaner interface {
	clean() error
}

func (_ local) clean() error {
	ref, err := getContainerRef("8080")
	_, err = local{}.run(exec.Command("docker", "rm", "-f", ref))
	return err
}

func (_ remote) deploy(name string) error {
	pretty.Logln("[INFO] deploying remotely...")
	_, err := remote{}.run(exec.Command("sudo", "docker", "run", "-d", "-p", "80:8080", name))
	switch err {
	case portAlreadyAllocated:
		pretty.Logln("[INFO] looks like port is already allocated. cleaning...")
		if err := remoteClean(); err != nil {
			return err
		}
		return remote{}.deploy(name)
	case nil:
		return nil
	}
	return err
}

func remoteClean() error {
	ref, err := getContainerRefGce("8080")
	_, err = remote{}.run(exec.Command("sudo", "docker", "rm", "-f", ref))
	return err
}

func getContainerRef(port string) (string, error) {
	stdout, err := local{}.run(exec.Command("docker", "ps", "-f", "publish="+port, "-q"))
	if err != nil {
		return "", err
	}
	return strings.Trim(stdout, "\n"), nil
}

func getContainerRefGce(port string) (string, error) {
	stdout, err := remote{}.run(exec.Command("sudo", "docker", "ps", "-f", "publish="+port, "-q"))
	if err != nil {
		return "", err
	}
	return strings.Trim(stdout, "\n"), nil
}
