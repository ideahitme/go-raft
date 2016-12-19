package cluster

import "math/rand"

func (c *Cluster) getClusterSize() uint16 {
	return uint16(len(c.cluster))
}

func (c *Cluster) minAcceptSize() uint16 {
	return c.getClusterSize()/2 + 1
}

func getElectionTimeout() int32 {
	r := rand.Int31n(150) + 150
	return r
}
