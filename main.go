package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ideahitme/go-raft/cluster"
	"github.com/ideahitme/go-raft/listener"
	"github.com/ideahitme/go-raft/pkg"
	"github.com/ideahitme/go-raft/server"
)

func main() {
	log.Infoln("Starting API server...")

	heartbeatChannel := make(chan *pkg.Heartbeat)
	electionChannel := make(chan *pkg.VoteRequest)
	logChannel := make(chan *pkg.Log)

	listen := &listener.Listener{
		HeartbeatChannel: heartbeatChannel,
		ElectionChannel:  electionChannel,
		LogChannel:       logChannel,
	}

	cluster := cluster.NewCluster([]*pkg.Node{}, nil, listen)
	go cluster.Run()

	server.StartAPIServer(listen)

}
