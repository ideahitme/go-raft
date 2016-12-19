package cluster

import "github.com/ideahitme/go-raft/pkg"

// HTTPClient client for interaction with the cluster
type HTTPClient struct {
}

// NewHTTPClient returns a new instance of a HTTPClient
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{}
}

// SendVoteRequest sends a request to a particular node and wait for response
// returns true if the candidate accepts the node
// if timeout or rejected returns false
func (c *HTTPClient) SendVoteRequest(n *pkg.Node) bool {
	return true
}

// SendHeartBeat sends a request to all nodes indiciating leader is alive
func (c *HTTPClient) SendHeartBeat() bool {
	return true
}

// SendLogUpdate sends a log update
func (c *HTTPClient) SendLogUpdate() bool {
	return true
}
