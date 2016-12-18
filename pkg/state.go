package pkg

type State struct {
	node   *Node
	leader *Node
}

func NewState() *State {
	n := NewNode(IP("192.168.0.1"))
	return &State{
		node: n,
	}
}

func (s *State) becomeLeader() {
	s.node.state = LeaderState
	s.leader = s.node
}

func (s *State) currentLeader() *Node {
	return s.leader
}
