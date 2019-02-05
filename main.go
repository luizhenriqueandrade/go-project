package main

import(
	"github.com/luizhenriqueandrade/go-project/api"
	"log"
	"net/http"
)

func main(){
	addr := "127.0.0.1:8081"
	router := http.NewServeMux()
	router.Handle("/", &api.HomepageHandler{})
	router.Handle("/new", &api.MovieHandler{})
	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}