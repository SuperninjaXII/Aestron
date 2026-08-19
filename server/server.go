package main

import (
	"fmt"
	"net/http"

	"github.com/SuperninjaXII/Aestron/components"
	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("src"))

	mux.Handle("/src/", http.StripPrefix("/src/", fileServer))
	mux.Handle("/", templ.Handler(components.Home()))
	fmt.Println("runing on port 8080")
	http.ListenAndServe(":8080", mux)
}
