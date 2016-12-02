package main

import (
	"fmt"
	"net/http"

	"github.com/tattsun/mytasks/storage"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func root(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func main() {
	storage.NewDB("/tmp/mytasks.db")

	goji.Get("/", root)
	goji.Serve()
}
