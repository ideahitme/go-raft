package pkg

import "math/rand"

// Leader current leader of the cluster
var Leader *Node
var Cluster []*Node

func getClusterSize() uint16 {
	return uint16(len(Cluster))
}

func minAcceptSize() uint16 {
	return getClusterSize()/2 + 1
}

func sendVoteRequest(n *Node) bool {
	//send http request to the Node
	return true
}

func getElectionTimeout() int32 {
	r := rand.Int31n(150) + 150
	return r
}

// RequestVote candidate request for votes to become a leader
func RequestVote() {
	reqChannel := make(chan bool)

	for _, node := range Cluster {
		node := node
		go func() {
			reqChannel <- sendVoteRequest(node)
		}()
	}

	var accepted uint16 = 1 // Node votes for itself
	var i uint16
	for i = 0; i < getClusterSize(); i++ {
		if accept := <-reqChannel; accept {
			accepted++
		}
	}

	if accepted >= minAcceptSize() {
		if isLeader := ClaimLeadership(); isLeader {

		}
	}
}

// ClaimLeadership send request to all nodes to claim a leadership
func ClaimLeadership() bool {
	claimChannel := make(chan bool)

	for _, node := range Cluster {
		node := node
		go func() {
			claimChannel <- sendVoteRequest(node)
		}()
	}

	var accepted uint16 = 1 // Node votes for itself
	var i uint16
	for i = 0; i < getClusterSize(); i++ {
		if accept := <-claimChannel; accept {
			accepted++
		}
	}

	if accepted >= minAcceptSize() {
		return true
	}
	return false
}
