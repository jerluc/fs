package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var portNo string
	var rootDir string
	flag.StringVar(&portNo, "port", "8000", "Port number to bind server")
	flag.StringVar(&rootDir, "root", ".", "Root directory to serve")
	flag.Parse()
	log.Fatal(http.ListenAndServe(":" + portNo, http.FileServer(http.Dir(rootDir))))
}
