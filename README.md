# Raft Consensus Algorithm
Raft consensus algorithm is an algorithms which allows to have a fault-tolerant replicated system (e.g. districtured hash table).  It is a widely used algorithm, 
which is proven to be simple yet efficient. 

The problems arises, when we have a set of machines (nodes), which should have a replicated state across the cluster. We have 
a client interacting with one of the machines (leader), and sending a request to either `set` or `get` a value. 

## State of a node

1. Follower - a slave node 
2. Candidate - node ready to become a leader
3. Leader - current master of the cluster

## Prinicipal of work

### Leader election

All nodes start in the `follower` state. There are two timeouts:

1. `Election` timeout - amount of time a follower waits until becoming a candidate. Each node gets a random value between `150ms` to `300ms`. 
2. ``

After the `election` timeout a `follower` become a `candidate` and starts a new 
`election term`. 