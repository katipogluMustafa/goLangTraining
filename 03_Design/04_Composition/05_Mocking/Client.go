/*
 * Sample program to show how you can personally mock concrete types
 * when you need to for your own packages or tests.
 */

package main

import "goLangTraining/03_Design/04_Composition/05_Mocking/pubsub"

type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

// mock is a concrete type to help support the mocking of the pubsub package.
type mock struct{}

// Publish and Subscribe make mock implement publisher interface
func (m *mock) Publish(key string, v interface{}) error {
	// Add your mock for the publish call
	return nil
}

func (m *mock) Subscribe(key string) error {
	// Add your mock for the subscribe call
	return nil
}

func main() {
	pubs := []publisher{
		pubsub.New("localhost"),
		&mock{},
	}
	//see how the publisher interface provides the level of decoupling the user needs.
	for _, p := range pubs {
		p.Publish("key", "value")
		p.Subscribe("key")
	}
}
