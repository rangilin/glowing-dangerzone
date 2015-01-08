package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Run file server on specified directory
func RunFileServer(dir string) {
	flagSet := flag.NewFlagSet("serve", flag.ContinueOnError)
	port := flagSet.Int("port", 80, "port to serve file server")
	flagSet.Parse(os.Args[2:])

	server := http.FileServer(http.Dir(dir))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), server))
}
