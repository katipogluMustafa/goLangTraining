package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 Expired")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped")
	}

}

/*
 * If you just wanted to wait, you could have used time.Sleep. One reason a timer may be useful is that you can cancel the timer before it expires.
 *The first timer will expire ~2s after we start the program, but the second should be stopped before it has a chance to expire.
 */
