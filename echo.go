package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

// Default port
var Port = "9999"

func init() {
	// https://gobyexample.com/command-line-flags && http://golang.org/pkg/flag/
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "A simple HTTP echo server that will echo the incoming request as the body of the response\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.StringVar(&Port, "port", Port, "port to listen on")
	flag.StringVar(&Port, "p", Port, "port to listen on (shorthand)")
}

// EchoHandler echos back the request as a response
func EchoHandler(writer http.ResponseWriter, request *http.Request) {

	log.Println("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")
	reqDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(reqDump))

	// let everyone in
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	// allow pre-flight headers
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	// request.Write(writer) // probably the same as the dump
	fmt.Fprintf(writer, "%s\n", reqDump)
}

func main() {

	flag.Parse()
	log.Println("Listening on port: " + Port)

	http.HandleFunc("/", EchoHandler)
	http.ListenAndServe(":"+Port, nil)
}
