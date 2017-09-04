package main

import (
	"errors"
	"log"
	"os/exec"
)

var resourceAlreadyProvisioned = errors.New("resource already provisioned")

func provision() error {
	log.Println("[INFO] provisioning...")
	if err := provisionNetwork(); err != nil {
		return err
	}
	if err := provisionInstance(); err != nil {
		return err
	}
	return nil
}

func provisionNetwork() error {
	log.Println("[INFO] provisioning network...")
	cmd := exec.Command("aws", "ec2", "create-vpc", "--cidr-block", "10.0.0.0/28", "--query", "'VpcVpcId'", "--output", "text")
	if err := execute(cmd); err != nil {
		return err
	}
	return nil
}

func provisionInstance() error {
	log.Println("[INFO] provisioning network...")
	return nil
}
