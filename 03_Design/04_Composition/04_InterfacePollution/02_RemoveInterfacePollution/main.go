/*
 * This is an example that removes the interface pollution by
 * removing the interface and using the concrete type directly.
 */
package main

// Server is our Server implementation.
type Server struct {
	host string

	// Pretend there are more...
}

func NewServer(host string) *Server {
	return &Server{host}
}

// Start allows the server to begin to accept requests.
func (s *Server) Start() error {
	// Pretend there is a specific implementation.
	return nil
}

// Wait prevents the server from accepting new connections.
func (s *Server) Wait() error {
	// Pretend there is a specific implementation.
	return nil
}

// Stop shuts the server down.
func (s *Server) Stop() error {
	// Pretend there is a specific implementation.
	return nil
}
