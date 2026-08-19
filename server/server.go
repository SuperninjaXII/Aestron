package main

import (
	"fmt"
	"net/http"

	"github.com/SuperninjaXII/Aestron/components"
	"github.com/a-h/templ"
	"github.com/aarol/reload"
)

func main() {
	var handler http.Handler = http.DefaultServeMux
	reloader := reload.New("./src")
	fileServer := http.FileServer(http.Dir("src"))

	http.Handle("/src/", http.StripPrefix("/src/", fileServer))
	http.Handle("/", templ.Handler(components.Home()))
	fmt.Println("runing on port 8080")
	handler = reloader.Handle(handler)
	http.ListenAndServe(":8080", handler)
}
