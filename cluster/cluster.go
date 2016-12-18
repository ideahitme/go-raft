package cluster

import "math/rand"
import "github.com/ideahitme/go-raft/server"

var client server.Client

func init() {
	client := server.NewClient()
}

func (s *State) getClusterSize() uint16 {
	return uint16(len(s.nodes))
}

func (s *State) minAcceptSize() uint16 {
	return s.getClusterSize()/2 + 1
}

func getElectionTimeout() int32 {
	r := rand.Int31n(150) + 150
	return r
}

// RequestVote candidate request for votes to become a leader
func (s *State) RequestVote() {
	reqChannel := make(chan bool)

	for _, node := range s.nodes {
		node := node
		go func() {
			reqChannel <- sendVoteRequest(node)
		}()
	}

	var accepted uint16 = 1 // Node votes for itself
	var i uint16
	for i = 0; i < s.getClusterSize(); i++ {
		if accept := <-reqChannel; accept {
			accepted++
		}
	}

	if accepted >= s.minAcceptSize() {
		if isLeader := s.ClaimLeadership(); isLeader {

		}
	}
}

// ClaimLeadership send request to all nodes to claim a leadership
func (s *State) ClaimLeadership() bool {
	claimChannel := make(chan bool)

	for _, node := range s.nodes {
		node := node
		go func() {
			claimChannel <- sendVoteRequest(node)
		}()
	}

	var accepted uint16 = 1 // Node votes for itself
	var i uint16
	for ; i < s.getClusterSize(); i++ {
		if accept := <-claimChannel; accept {
			accepted++
		}
	}

	if accepted >= s.minAcceptSize() {
		return true
	}
	return false
}
