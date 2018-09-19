* Remember problems are solved in the concrete not in the abstract, so we need concrete implementations, solutions first in order to know how to decouple.This is gonna simplify your code and allow you focus on what's important which is the problem and remember problem is the data.
* How do you know when you done a piece of code ? 
    * For me done comes with two different answers
        * The first one is, you really have gotta make sure that you have unit tests.You've gotta have some level of test coverage.For me overall I want to see 70-80% test coverage before you can come to me and say you're done.
        * The other part of the answer of this question, is the code decoupled from the change we expect to happen ? 
            * Remember you're writing code for today, design and architect for tomorrow.
            * So deciding to decouple may or may not happen immediately but i wanna ask the question, do we know what has to be decoupled and do we wanna do that ? I can ask this question because decoupling is part 2 of everything I do, it is refactoring.We saw problems in the concrete first, we refactor into the decoupling.
---
 ![Layered API Design](img/layeredAPIDesign.png)
* Code that is testable usually means the data that we're passing in is reproducible on data coming out.
* Remember! 
    * Testability is about the data, the code we're writing is about the data, decoupling is about the data, the problem you're trying to solve is about the data.             
* As we start write more code, you're gonna see me focus on
    * What are the real problems in front of us
    * How do we layer this API, not only so usable but testable.It doesn't mean we need interfaces, we need strong inputs and outputs around the data that we can test against.
    * How do we refactor code to then deal with change once we have concrete solution.  
---
* The idea of concrete implementation
    * Could be a prototype that we can even put it production right away.
        * The idea is we have to be able to get code into production faster, the code that has integrity, code that we can really start solving problems because for me, when we work on a piece of code, it never leaves our laptop, it never gets into production.We have to be better at all of this.
---

* Why would I say and believe that a function is always gonna be more precise than a method can ever be ?
    * This is true, functions can always be more precise.
    * The reason a function can always be more precise is because a function requires you to pass in all of the input in order for it perform data transformation.Methods when they are not designed properly hide information, and when we start designing method based APIs, we still have to take extra care so that we make sure we are not hiding anything.
        * When we hide things, we are setting up fraud and misuse.        
 * Imagine I needed to send an email to a user
 
 ```go

type user struct {
	    name    string
	    email   string
	    // MORE FIELDS
}

func (u *User) SendEmail(){
	//
}
```
* I already have the user state, then add the behaviour to the data.But this is a HORRIFIC API design.This is exactly the opposite of what I want you to do.This API is not precise it is riddled with fraud because this API right now doesn't tell us what that user value needs to have in terms of initialization for this method to not the fail and we don't want things fail at runtime because the state of a value is incomplete.That is a major violation.**DO NOTE DO** things like this.Because the only time we're gonna know if there is a problem is in production.
    * Today maybe send email is using name and email but what if I change SendEmail() one day to require to use of user's age?
        * Nobody calling send email not even the compiler knows now we need age to be set. 
        * Which means that Now anybody who send email will fail on sending email because they didn't know age had to be initialized. We don't know this until production.
    * By the way, functions like this no better than methods.
        ```go
        func SendEmail(u *User)
        ``` 
* So how do we write an API that's precise using this example ? 
    * I wanna see a function that says give me a name and give me a email.
    * It doesn't get any more precise than this. 
```go
func SendEmail(name string, email string)
```       
* I am sorry but if we need age tomorrow I'm going to absolutely take the cost of the API breaking.
    * Most of us work on private code base, it is okay to break APIs.
        * I demand that the APIs break because now the compiler can tell us everywhere in the codebase we were going to have a problem with SendEmail() before it makes into production.
    * I do have code in production level software that has functions that can take up to 8,10,12 parameters.
        * The more precise is the better.
        * So don't get into the idea we need the abstract the input.
        * Because abstracting the input hurt your testing, debugging and users.    
```go
func SendEmail(name string, email string, age int)
```       
* This is why I am going with function based APIs first.