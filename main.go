package main

import (
	"flag"
	"log"

	"github.com/kr/pretty"
)

var verbose = flag.Bool("verbose", false, "verbosity flag")

type temp interface {
	build() error
	test() error
	provision() error
	configure() error
	deploy() error
	monitor() error
}

const name = "atec/www"

func main() {

	flag.Parse()

	_, err := build(name)
	if err != nil {
		pretty.Logln("[FATAL] failed to build")
		log.Fatal(err)
	}
	if err := (remote{}.deploy(name)); err != nil {
		pretty.Logln("[FATAL] failed to deploy")
		log.Fatal(err)
	}
	pretty.Logln("[INFO] success!")
}
