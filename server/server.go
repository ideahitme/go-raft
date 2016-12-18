package server

import (
	"net/http"
)

// StartAPIServer starts the api server for interaction with the cluster
func StartAPIServer() {
	http.HandleFunc("/request-vote", onRequestVote)
	http.ListenAndServe(":3000", nil)
}

func onRequestVote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello sir"))
}
