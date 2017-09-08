package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

type executive interface {
	execute(cmd *exec.Cmd) (string, error)
}

type local struct{}

func (_ local) execute(cmd *exec.Cmd) (string, error) {

	pretty.Logf("[INFO] running %s...", cmd.Args)
	stdout, err := cmd.Output()
	if *verbose {
		pretty.Log("[DEBUG] stdout:", string(stdout))
	}

	// cmd usually returns exit error if not nil, will panic otherwise
	if err != nil {
		pretty.Logln("[INFO] checking exit error...")
		return string(stdout), checkErr(err.(*exec.ExitError))
	}

	return string(stdout), nil
}

type remote struct{}

func (_ remote) execute(remoteCmd *exec.Cmd) (string, error) {
	args := strings.Join(remoteCmd.Args, " ")
	return local{}.execute(exec.Command("gcloud", "compute", "ssh", "atec", "--zone", "us-east1-b", "--command", args))
}

var (
	portAlreadyAllocated       = errors.New("port is already allocated") // docker
	resourceAlreadyProvisioned = errors.New("resource already exists")   // gce
	tmpDirDoesntExist          = errors.New("tmp dir doesn't exist")     // unix fs
)

func checkErr(err *exec.ExitError) error {
	pretty.Log("[INFO] stderr:", string(err.Stderr))
	if bytes.Contains(err.Stderr, []byte("port is already allocated")) {
		return portAlreadyAllocated
	}
	if bytes.Contains(err.Stderr, []byte("already exists")) {
		return resourceAlreadyProvisioned
	}
	return err
}
