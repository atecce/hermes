package main

import (
	"os/exec"

	"github.com/kr/pretty"
)

type deployer interface {
	deploy(name string) error
}

func (_ local) deploy(name string) error {
	pretty.Logln("[INFO] deploying locally...")
	_, err := local{}.run(exec.Command("vagrant", "share", "--http", "8080"))
	return err
}

type cleaner interface {
	clean() error
}
