package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"amtsim/internal/amt"
	"amtsim/internal/cira"
	"amtsim/internal/logging"
	"amtsim/internal/server"
)

func main() {
	var (
		listenAddr string
		mpsAddr    string
		mpsUser    string
		mpsPass    string
		amtUuid    string
		debugMode  bool
	)

	flag.StringVar(&listenAddr, "listen", ":16992", "Address to listen on")
	flag.StringVar(&mpsAddr, "mps", "", "MPS Address (host:port) to connect to (CIRA)")
	flag.StringVar(&mpsUser, "mps-user", "", "MPS Username")
	flag.StringVar(&mpsPass, "mps-pass", "", "MPS Password")
	flag.StringVar(&amtUuid, "uuid", "1d8aa800-918d-11e4-b33f-123400000000", "AMT UUID")
	flag.BoolVar(&debugMode, "debug", false, "Enable debug logging")
	flag.Parse()

	if debugMode {
		logging.SetLevel(logging.LEVEL_DEBUG)
	}

	// Try to load saved credentials from a prior sesson.
	if savedCreds, err := cira.LoadCredentials(); err == nil && savedCreds != nil {
		logging.Info("Loaded saved CIRA credentials for user: %s", savedCreds.Username)
		mpsUser = savedCreds.Username
		mpsPass = savedCreds.Password
	}

	srv := server.NewServer(listenAddr)
	// Initialize UUID
	amt.SetUUID(amtUuid)
	srv.Start()

	ciraClient := cira.NewClient(mpsAddr, mpsUser, mpsPass, amtUuid)
	ciraClient.Start()

	logging.Info("Main: Wait for terminate...")

	// Wait for interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	logging.Info("Main: Shutting down...")
}
