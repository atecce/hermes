package main

import (
	"log"

	"github.com/kr/pretty"
)

type temp interface {
	build() error
	test() error
	provision() error
	configure() error
	deploy() error
	monitor() error
}

func main() {
	ref, err := build()
	if err != nil {
		pretty.Logln("[FATAL] failed to build")
		log.Fatal(err)
	}
	_, err = deploy(ref)
	if err != nil {
		pretty.Logln("[FATAL] failed to deploy")
		log.Fatal(err)
	}
}
