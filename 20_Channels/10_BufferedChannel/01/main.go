package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*
By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.
*/

/*
 * BUFFERED CHANNELS = QUEUES
 * Buffered channels are first-in first-out (FIFO) queues of bounded capacity. The capacity is fixed at the time of creation of the queue – queues cannot be resized on the fly.
 */
