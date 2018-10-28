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

## Behaviour As Context

* If you custom error type can have any one of these 4 methods, it doesn't have to have all of them just one, then I want the custom error type to be defined as unexported with unexported fields because if it is unexported with unexported fields, you're forcing that your user can never ever go from decoupled state to that concrete right ? , we can't type assert to the concrete , you're really helping them even though they might feel like they're being restricted.

* Here The 4 Methods:
    1. Temporary
    2. Timeout 
    3. NotFound
    4. NotAuthorized      

* If you can add anyone of these 4 we can maintain decoupling with custom error types by making the custom error type unexported.    

* Remember, Temporary() covers tremendous amount.Temporary is like blanket statement that you have an integrity issue or you don't.
* If people complain that you've made the custom error type unexported and you have valid Temporary() method just tell them look we're not going to export the type, even if you feel like you need to inspect it directly yourself.Let's fix Temporary() to be more accurate.
    * What's brilliant is that we can fix Temporary() to be more accurate without creating a cascading code change.Because this code doesn't change if we change the implementation of the Temporary().
    
* If you have a custom error type like the json package did, timeout, temporary, notfound, notauthorized, they don't work for those types.Those types have to be exported.These may need to move from decoupling back to the concrete.It's a scary situation but that is what it is. Most of the time that temporary method is going to work for you.

* So I don't have a problem with custom error types as soon as the error variables don't work.But I'd love to make sure that those custom error types unexported with unexported fields and that we apply method sets of behaviour that the user can bind to to maintain their error handling from decoupling state.   

## Wrapping Errors

* You can't separate error handling and logging, these are just one thing that we've gotta bring them together if you want any consistency.

* As an insurance policy, to be able to find bugs, when errors occur we log them. Reality is that there is too much activity on our systems today, our user bases grow almost million user almost overnight, so logging that much has huge significant cost.A lot of times logging is going to create large amount of allocations which is going to put a lot of pressure on your heap(that's not unique to go but we're talking about go). 

* So I want you to consider that logging is important but we've got to constantly balance signal to noise in the log because if you're writing data to your logs and you end up never ever reading or using then you're wasting CPU cycles on something that you could've been doing actual real work and it just go beyond the CPU cycles of your process, you're eating network bandwidth, diskIO bandwidth, other complexities that go through the entire system.
    * During development I really wanna make sure that we always have good level of signal in our logs, and we're logging from trace perspective which minimum we need but then we're logging the errors in a way that there is always enough context if we take the time or need take time to look at it.

How do you make sure there is enough context in the log both from tracing perspective bare minimum and in error perspective not duplicate errors throughout a log and at the same time minimize those logs and then have a consistent pattern that we all can follow and review during code reviews?             

* What do I mean when i am talking about error handling for an error ?
    * When a piece of code decides to handle an error, it means that this piece of code responsible for logging it, logging the full context of it.It also means that the code has to make decision, can we recover or not ?, if the answer is no then error handling means we shut down the app either with the stack trace on the panic call or with os.Exit() , If we can recover, then that code has to recover the application back to its correct state and keep it going.
    * At the end of the day when that error returns, it never ever returns error back up, error stops there, either the app recovered or it shuts down. But we've also logged the error.

[ Wrapping Errors](/04_ErrorHandling/06_WrappingErrors/main.go)    
      