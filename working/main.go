package main

import (
	"log"
	"net/http"
	"os"

)

func main() {
	l := log.New(os.Stdout, "product-apt", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)
}