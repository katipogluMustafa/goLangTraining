package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

/* Using Nilakantha Terms, Pi Calculation*/

/* We're gonna create the virtual scoreboard by using ANSI escape codes.*/
/* Deprecated since the windows version doesn't respond somehow(find solution later for this)
const ANSIClearScreenSequence = "\033[H\033[2J"
const ANSFirstSlotScreenSequence = "\033[2;0H"
const ANSISecondSlotScreenSequence = "\033[3;0H"
*/
/* pichan used the update the current value of the pi on the scoreboard */
var pichan = make(chan float64)

var computationDone = make(chan bool, 1)

// Number of Nilakantha terms for the scoreboard
var termsCount int

/* Virtual Scoreboard*/
func printCalculationSummary() {
	fmt.Print( /* ANSIClearScreenSequence */ )
	fmt.Println( /* ANSFirstSlotScreenSequence ,*/ "Computed value of Pi:\t\t", <-pichan)
	fmt.Println( /* ANSISecondSlotScreenSequence,*/ "# of Nilakantha Terms:\t\t", termsCount)
}

func pi(n int) float64 {
	ch := make(chan float64)
	// Calculate each Nilakantha term in its own goroutine
	for k := 1; k <= n; k++ {
		go nilakanthaTerm(ch, float64(k))
	}
	f := 3.0

	// Sump up the calculated Nilakantha terms for n steps
	for k := 1; k <= n; k++ {
		termsCount++
		f += <-ch
		pichan <- f
	}

	// notify that the computation is done over the channel
	computationDone <- true
	return f
}

func nilakanthaTerm(ch chan float64, k float64) {
	j := 2 * k
	if int64(k)%2 == 1 {
		ch <- 4.0 / (j * (j + 1) * (j + 2))
	} else {
		ch <- -4.0 / (j * (j + 1) * (j + 2))
	}
}

func main() {
	ticker := time.NewTicker(time.Millisecond * 108)

	/*
	 * If the user wants to prematurely break out of the program by issuing a Ctrl+C
	 * 	 we tell the signal package to notify us over this interrupt channel
	 */
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go pi(5000)

	/* update values on the scoreboard */
	go func() {
		for range ticker.C {
			printCalculationSummary()
		}
	}()

	for {
		select {

		// Handle the case when the computation has ended,
		// we can end the program(exit out of this infinite loop)
		case <-computationDone:
			ticker.Stop()
			fmt.Println("Program done calculating Pi.")
			os.Exit(0)

		// If the user interrupts the program, we end the loop
		case <-interrupt:
			ticker.Stop()
			fmt.Println("Program interrupted by the user.")
			return

		}
	}
}
