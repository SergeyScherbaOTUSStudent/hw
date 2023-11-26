package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	envdirr := args[0]
	cmd := args[1:]
	env, err := ReadDir(envdirr)
	if err != nil {
		log.Fatalf("can't read environment dir %s", envdirr)
	}
	os.Exit(RunCmd(cmd, env))
}
