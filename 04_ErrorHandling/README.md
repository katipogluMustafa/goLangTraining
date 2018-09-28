# Error Handling

## Default Error Values

In go errors just values that can be anything you want.

From API design perspective, error handling is about one thing.
    
* It's about showing respect to the user of your API.Giving your user enough context to make an informed decision about state of the application and giving them enough information to be able to either recover or shut down.
    
If your application loses integrity, you have a responsibility to shut it down.

* There is two ways to shut down an application in go. You'll chose one over the other depending if you need stack trace or not. If you need stack trace you're gonna call panic, if you don't you just call os.Exit().
    * You can call [os.Exit()](https://golang.org/pkg/os/#Exit), that's the fastest way.
    * Or you can call the built-in function panic.
---

* Use the if statement to handle your negative path logic and keep your positive path logic on a straight line of sight.

```go
package main

import (
	"fmt"
)

func main(){
    if err := webCall(); err != nil {
    	fmt.Println(err)
    	return
    }	
    
    fmt.Println("Life is good:P")
}

// webCall() performs a web operation.
func webCall() error{
	
	// code for webCall()....
	
	return nil
}
```     

* I really want you to avoid else blocks for if.
    * Else causes code an order of magnitude more complicated to read.

* Anytime I see an error inside if , if error  not equal to nil, what it really saying is, is there a value where a piece of data stored concrete inside of the error interface.If there is we have an error.     
* What if web call could return multiple errors ? 
    * At that point now we have to deal with different mechanism to give the user enough context to make an informed decision.
    
## Error Variables

I want this error variables which used in multiple places at the top of the source code file where they're used.

* What happens if error variable doesn't give the user enough context ?
    * When that happens , next thing we're allowed to do is create our own custom error types.But I do not wanna be polluting our applications with custom error types just for the sake of using them.We wanna use errorString type and variables first until we no longer get enough context from it.  

## Types As Context

* Be careful when we're using empty interface.
    * Don't use it to define generic APIs.
    * We should be using empty interface when we have to pass data around where that data can be hidden without problems.
     * or [in this case](/04_ErrorHandling/03_TypeAsContext/main.go) where we're gonna use the reflect package because what we wanna do is at run-time or dynamically inspect the concrete data.
        * This is great if you wanna do model validation. reflect package is fantastic for model validation or like we're doing here Unmarshaling.
        
* We wanna maintain error handling from decoupled state.
    * Because once we switched from decoupled stated to concrete then any improvements we make to error handling against those concrete types could cause a cascading effect of change throughout the code base. 

* Type as context can be powerful when you need to move concrete data across program boundaries where both sides need to work with the concrete data itself then this idea of type as context can come in really really handy to move data across program boundaries maintaining levels of decoupling.            