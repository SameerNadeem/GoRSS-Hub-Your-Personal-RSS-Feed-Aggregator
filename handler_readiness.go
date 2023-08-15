package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, 200, struct{}{})
}
