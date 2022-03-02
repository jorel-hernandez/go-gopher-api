package main

import (
	"flag"
	"fmt"
	gopher "go-gopher-api/gopherapi/cmd/gopherapi/pkg"
	"go-gopher-api/gopherapi/cmd/gopherapi/pkg/server"
	"go-gopher-api/gopherapi/cmd/gopherapi/pkg/storage/inmem"
	"go-gopher-api/gopherapi/cmd/sample-data"
	"log"
	"net/http"
)

func main() {
	withData := flag.Bool("withData", false, "initialize the api with some gophers")

	flag.Parse()

	var gophers map[string]*gopher.Gopher
	if *withData {
		gophers = sample.Gophers
		fmt.Println("with data")
	}

	repo := inmem.NewGopherRepository(gophers)
	s := server.New(repo)

	fmt.Println("The gopher server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
