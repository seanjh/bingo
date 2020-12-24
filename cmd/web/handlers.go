package main

import (
	"net/http"
)

func base(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	return
}
