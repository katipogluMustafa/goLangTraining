package pubsub

type PubSub struct {
	host string

	// Pretend there are more...
}

// New creates a PubSub value for use.
func New(host string) *PubSub {
	ps := PubSub{
		host: host,
	}

	// ...

	return &ps
}

// Publish sends the data for the specified key.
func (ps *PubSub) Publish(key string, v interface{}) error {
	//...
	return nil
}

// Subscribe sets up an request to receive messages for the specified key.
func (ps *PubSub) Subscribe(key string) error {
	//...
	return nil
}
