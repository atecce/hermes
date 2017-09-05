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
	err := build()
	if err != nil {
		pretty.Logln("[FATAL] failed to build")
		log.Fatal(err)
	}
	err = configure()
	if err != nil {
		pretty.Logln("[FATAL] failed to configure")
		log.Fatal(err)
	}
	err = deploy()
	if err != nil {
		pretty.Logln("[FATAL] failed to deploy")
		log.Fatal(err)
	}
}
