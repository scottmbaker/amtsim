package server

import (
	"log"
	"net/http"

	"amtsim/internal/auth"
	"amtsim/internal/logging"
	"amtsim/internal/wsman"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Start() {
	logging.Info("Starting AMT Simulator on %s...", s.addr)

	// Start server in a goroutine
	go func() {
		if err := s.Run(); err != nil {
			logging.Fatal("Server failed: %v", err)
		}
	}()
}

func (s *Server) Run() error {
	mux := http.NewServeMux()

	// Auth Middleware
	authMiddleware := auth.NewDigestMiddleware("Digest:A3829B3827DE4D33D4449B366831FD01", map[string]string{
		"admin": "password", // TODO: why is this hardcoded?
	})

	// Wrap the handler
	protectedHandler := authMiddleware.Middleware(s.handleWSMAN)
	mux.HandleFunc("/wsman", protectedHandler)

	log.Printf("Listening on %s", s.addr)
	return http.ListenAndServe(s.addr, mux)
}

func (s *Server) handleWSMAN(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Dispatch to WSMAN handler
	response, err := wsman.Handle(r)
	if err != nil {
		log.Printf("WSMAN Handle error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/soap+xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(response); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
