package server

// Client client for interaction with the cluster
type Client struct {
}

// NewClient returns a new instance of a client
func NewClient() *Client {
	return &Client{}
}

func (c *Client) sendVoteRequest() bool {
	return true
}
