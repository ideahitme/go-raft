package cluster

import "github.com/ideahitme/go-raft/pkg"
import "github.com/ideahitme/go-raft/listener"
import "fmt"
import "time"

type Cluster struct {
	cluster  []*pkg.Node
	node     *pkg.Node
	client   *HTTPClient
	listener *listener.Listener
}

func NewCluster(nodes []*pkg.Node, node *pkg.Node, listener *listener.Listener) *Cluster {
	client := NewHTTPClient()
	return &Cluster{
		cluster:  nodes,
		node:     node,
		client:   client,
		listener: listener,
	}
}

func (c *Cluster) Run() {
	select {
	case <-c.listener.HeartbeatChannel:
		fmt.Println("Got a heartbeat from the leader")
	case <-time.After(BootTime):
		fmt.Println("Time to become a leader")
	}
}

func (c *Cluster) ListenForHearbeat() {
	for {
		select {
		case <-c.listener.HeartbeatChannel:

		case <-time.After(HeartbeatTimeout):

			break
		}
	}
}

// RequestVote candidate request for votes to become a leader
func (c *Cluster) RequestVote() {
	reqChannel := make(chan bool)

	for _, node := range c.cluster {
		node := node
		go func() {
			reqChannel <- c.client.SendVoteRequest(node)
		}()
	}

	var accepted uint16 = 1 // Node votes for itself
	var i uint16
	for i = 0; i < c.getClusterSize(); i++ {
		if accept := <-reqChannel; accept {
			accepted++
		}
	}

	if accepted >= c.minAcceptSize() {
		if isLeader := c.ClaimLeadership(); isLeader {

		}
	}
}

// ClaimLeadership send request to all nodes to claim a leadership
func (c *Cluster) ClaimLeadership() bool {
	claimChannel := make(chan bool)

	for _, node := range c.cluster {
		node := node
		go func() {
			claimChannel <- c.client.SendVoteRequest(node)
		}()
	}

	var accepted uint16 = 1 // Node votes for itself
	var i uint16
	for ; i < c.getClusterSize(); i++ {
		if accept := <-claimChannel; accept {
			accepted++
		}
	}

	if accepted >= c.minAcceptSize() {
		return true
	}
	return false
}
