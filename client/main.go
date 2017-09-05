package main

import (
	"log"

	"github.com/kr/pretty"
)

type temp interface {
	clone() error
	build() error
	test() error
	provision() error
	configure() error
	deploy() error
	monitor() error
}

func main() {
	err := clean()
	if err != nil {
		pretty.Logln("[FATAL] failed to clean")
		log.Fatal(err)
	}
	err = clone()
	if err != nil {
		pretty.Logln("[FATAL] failed to clone")
		log.Fatal(err)
	}
	err = build()
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
