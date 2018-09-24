package main

import (
	"log"
	"os"
)

/*
 * Sample program to show how to use the log package from the standard library.
 */

func init() {
	// Change the default output device from the stderr to stdout.
	log.SetOutput(os.Stdout)

	// Set the prefix string for each log line.
	log.SetPrefix("Trace: ")

	// Set the extra log info.
	setFlags()

}

// setFlags add extra information on each log line.
func setFlags() {
	/*
	   Ldate		   // the date: 2009/01/23
	   Ltime           // the time: 01:23:23
	   Lmicroseconds   // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	   Llongfile       // full file name and line number: /a/b/c/d.go:23
	   Lshortfile      // final file name element and line number: d.go:23. overrides Llongfile
	   LstdFlags       // Ldate | Ltime // initial values for the standard logger
	*/
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("Main function started")

	names := []string{"Mustafa", "Ahmet", "Yusuf", "Abdullah"}

	log.Printf("These are named %+v\n", names)

	log.Fatalln("Terminate Program")

	log.Fatalln("main function ended")
}
