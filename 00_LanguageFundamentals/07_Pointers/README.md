# Pointers

* Everything in go pass by value.

* The Stack is a data structure that every thread is given.
At the operating system level your stack is contiguous block of memory and usually it is allocated to be 1MB.1 MB of memory for every stack.
* Every Goroutine has 2kb of memory.Goroutines are out path of execution. 

* Pointers are literal types[unnamed types].

* You can not read and write to memory unless you understand memory's type. So you have to declare type of the pointer if you want to get the address of anything.
```
func increment(inc *int){ // Can't be (i *) since we need to know the type.
    *inc++ // Indirect memory read and write through pointer indirection.
}
```

* Memory below the active frame has no longer integrity because its going to be reused.