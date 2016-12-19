//Minimum layer to enable HTTP communication across the cluster
package server

import (
	"net/http"

	"github.com/ideahitme/go-raft/listener"
)

// StartAPIServer starts the api server for interaction with the cluster
func StartAPIServer(c *listener.Listener) {
	http.HandleFunc("/api/request-vote", c.OnRequestVote)
	http.HandleFunc("/api/heartbeat", c.OnHeartbeat)
	http.HandleFunc("/api/log-update", c.OnLogUpdate)
	http.ListenAndServe(":3000", nil)
}
