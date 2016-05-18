package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/nytimes/gziphandler"
)

func main() {
	var portNo string
	var rootDir string
	flag.StringVar(&portNo, "port", "8000", "Port number to bind server")
	flag.StringVar(&rootDir, "root", ".", "Root directory to serve")
	flag.Parse()
	handler := gziphandler.GzipHandler(http.FileServer(http.Dir(rootDir)))
	log.Fatal(http.ListenAndServe(":" + portNo, handler))
}
