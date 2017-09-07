package main

import (
	"os/exec"

	"github.com/kr/pretty"
)

func provision() error {
	pretty.Logln("[INFO] provisioning...")
	cmd := exec.Command("gcloud", "compute", "instances", "create", "atec", "--zone", "us-east1-b", "--tags", "http-server")
	_, err := execute(cmd)
	switch err {
	case resourceAlreadyProvisioned:
		pretty.Logln("[INFO] resource already provisioned. skipping")
		return nil
	default:
		return err
	}
}
