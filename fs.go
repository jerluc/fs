package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/rs/cors"
)

func main() {
	var portNo string
	var rootDir string
	flag.StringVar(&portNo, "port", "8000", "Port number to bind server")
	flag.StringVar(&rootDir, "root", ".", "Root directory to serve")
	flag.Parse()
	mux := http.NewServeMux()
	handler := gziphandler.GzipHandler(http.FileServer(http.Dir(rootDir)))
	mux.Handle("/", handler)
	wrapped := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":"+portNo, wrapped))
}
