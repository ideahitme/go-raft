package main

import (
	log "github.com/Sirupsen/logrus"
	_ "github.com/ideahitme/go-raft/pkg"
	"github.com/ideahitme/go-raft/server"
)

func main() {
	log.Infoln("Starting API server...")
	server.StartAPIServer()
}
