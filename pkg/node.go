package pkg

// NodeState state current state of a node
type NodeState uint8

// IP address of the node
type IP string

const (
	_                        = iota
	FollowerState  NodeState = iota
	CandidateState NodeState = iota
	LeaderState    NodeState = iota
)

// Node corresponds to the machine
type Node struct {
	State   NodeState
	Address IP
}

// NewNode create a new Node pointer
func NewNode(ip IP) *Node {
	return &Node{
		State:   FollowerState,
		Address: ip,
	}
}
