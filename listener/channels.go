package listener

import "net/http"
import "github.com/ideahitme/go-raft/pkg"

type Listener struct {
	HeartbeatChannel chan *pkg.Heartbeat
	ElectionChannel  chan *pkg.VoteRequest
	LogChannel       chan *pkg.Log
}

func (c *Listener) OnRequestVote(w http.ResponseWriter, r *http.Request) {
	c.ElectionChannel <- nil
}

func (c *Listener) OnHeartbeat(w http.ResponseWriter, r *http.Request) {
	c.HeartbeatChannel <- nil
}

func (c *Listener) OnLogUpdate(w http.ResponseWriter, r *http.Request) {
	c.LogChannel <- nil
}
