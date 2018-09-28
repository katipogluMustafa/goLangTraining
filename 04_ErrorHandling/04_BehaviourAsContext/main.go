// Shows how to implement behaviour as context

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// client represents a single connection in the room
type client struct {
	name   string
	reader *bufio.Reader
}

/*
 * TypeAsContext shows how to check multiple types of possible custom error
   types that can be returned from the net package.
*/
func (c *client) TypeAsContext() {
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			case *net.OpError:
				if !e.Temporary() {
					log.Fatalln("Temporary, client leaving chat.")
					return
				}
			case *net.AddrError:
				/*
				 * If there is a Temporary() method like this remember
				    errors just values in go they can be anything we need them to be both in state and behaviour
				 * This Temporary method is beautiful because if it is temporary
				    then we know that we're still in a state of integrity, just keep going.
				    If it's not temporary, we've lost integrity maybe that listeners gone down or socket is dropped.
				*/
				if !e.Temporary() {
					log.Println("Temporary, client leaving chat.")
					return
				}
			case *net.DNSConfigError:
				if !e.Temporary() {
					log.Println("Temporary, client leaving chat.")
					return
				}
			default:
				if err == io.EOF {
					log.Println("EOF: Client leaving chat.")
					return
				}
				log.Println("read-routine", err)
			}
		}
		fmt.Println(line)
	}
}

// temporary is declared to test for the existence of the method coming
// from the net package.
type temporary interface {
	Temporary() bool
}

// BehaviourAsContext shows how to check behaviour of an interface
// that can be return from the net package.
func (c *client) BehaviourAsContext() {
	line, err := c.reader.ReadString('\n')

	if err != nil {
		switch e := err.(type) {
		case temporary:
			if !e.Temporary() {
				log.Println("Temporary: Client leaving chat.")
				return
			}
		default:
			if err == io.EOF {
				log.Println("EOF: Client leaving chat.")
				return
			}
			log.Println("read-routine", err)
		}
	}
	fmt.Println(line)

}

func main() {
	// Not executable, see the method implementations.
}

/*
 * Btw, there is a naming convention for custom error types.
 * They end with the Error for example OpError, AddrError, DNSConfigError
 */
