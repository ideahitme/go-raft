# Go-raft

Simple implementation of the Raft Consensus Algorithm in Go. At this moment with the sole purpose of self-learning
and exploration

## Raft Consensus Algorithm
Raft consensus algorithm is an algorithm which allows to have a fault-tolerant replicated system (e.g. distributed hash table).  It is a widely used algorithm, 
which is proven to be simple yet efficient. 

The problems arises, when we have a set of machines (nodes), which should have a replicated state across the cluster. We have 
a client interacting with one of the machines (leader), and sending a request to either `set` or `get` a value. 

The following is the high-level description of the algorithm followed by the list of resources used

## State of a node

1. Follower - a slave node 
2. Candidate - node ready to become a leader
3. Leader - current master of the cluster

## Priniciples of work

### Leader election

All nodes start in the `follower` state. There are two timeouts:

1. `Election` timeout - amount of time a follower waits until becoming a candidate. 
Each node gets a random value between `150ms` to `300ms`. 
2. `Heartbeat` timeout - amount of time followers are waiting for pings from the leader. If timeout is exceeded, 
a node will strive to become a candidate.  

After the `election` timeout a `follower` become a `candidate` and starts a new 
`election term`. 

### Log replication

### Safety

## References

1. https://www.usenix.org/conference/atc14/technical-sessions/presentation/ongaro
2. https://raft.github.io/raft.pdf
3. https://raft.github.io/
4. http://thesecretlivesofdata.com/raft/

## License

MIT License

Copyright (c) 2016 Yerken (ideahitme)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
