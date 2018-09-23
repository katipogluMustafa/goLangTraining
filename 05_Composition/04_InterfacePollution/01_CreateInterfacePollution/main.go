/*
 * This is an example that creates interface pollution
 * by improperly using an interface when one is not needed.
 */
package main

// Server defines a contract for tcp servers.
type Server interface {
	Start() error
	Stop() error
	Wait() error
}

// server is our Server implementation
type server struct {
	host string

	// ...
}

/*
 * Factory functions initialize concrete type should return the concrete values or pointers
 * The caller will deal with decoupling if it is needed.
 */

func NewServer(host string) Server {

	return &server{host: host}
}

func (s *server) Start() error {
	// Pretend there is a specific implementation.
	return nil
}

func (s *server) Stop() error {
	// Pretend there is a specific implementation.
	return nil
}

func (s *server) Wait() error {
	// Pretend there is a specific implementation.
	return nil
}

func main() {

	// Create a new Server
	srv := NewServer("Localhost")

	// Use the API
	srv.Start()
	srv.Stop()
	srv.Wait()
}

// =============================================================================
/*
 * Smells:
 * The package declares an interface that matches the entire API of its own concrete type.
 * The interface is exported but the concrete type is unexported.
 * The factory function returns the interface value with the unexported concrete type value inside.
 * The interface can be removed and nothing changes for the user of the API.
 * The interface is not decoupling the API from change.
 */
