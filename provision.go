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

	vpcId, err := execute(exec.Command("aws", "ec2", "create-vpc", "--cidr-block", "10.0.0.0/28", "--query", "Vpc.VpcId", "--output", "text"))
	if err != nil {
		return err
	}

	_, err = execute(exec.Command("aws", "ec2", "modify-vpc-attribute", "--vpc-id", vpcId, "--enable-dns-support", "{\"Value\":true}"))
	if err != nil {
		return err
	}

	_, err = execute(exec.Command("aws", "ec2", "modify-vpc-attribute", "--vpc-id", vpcId, "--enable-dns-hostnames", "{\"Value\":true}"))
	if err != nil {
		return err
	}

	return nil
}

func provisionInstance() error {
	log.Println("[INFO] provisioning instance...")
	return nil
}
