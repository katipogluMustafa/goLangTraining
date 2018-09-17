## Semantic Usage
Remember, semantic consistancy is everything!

* The code, functions, the factory method, other methods, has to respect the chosen semantic
* If you're working with struct type that you are not familiar with,look for the factory function(generally right after the type) 
* Use semantic knowledge to understand how to work with data correctly
* When you're defining your own struct type 
    * If you are **not sure** what to use then we are gonna use **Pointer Semantics** 

### When to use Pointer Semantics ?

Not all the data can leverage value semantics especially when it comes to struct types.

* You have got to chose if you're defining your own struct type what semantic is gonna be in play.
    * If you are not sure what to use, then we are gonna use pointer semantics
    * If you absolutely sure we can use value semantics,then we want to use value semantics.
 

### When to use Value Semantics 

Value semantics are very important of your ability to keep the heap clean which is going to give us tremendous amount of performance.

* With Built-in Types {int, bool,string...}

* Fields in struct {See Exception 1}

* With Reference Types {See Exception 2}

### Exceptions:

**Exception - 1**: 

You can use Pointer semantics in struct fields instead of value semantics If you are dealing with sequal database and you have a struct that's going to be used to Marshal and Unmarshal data, the concept of null doesn't exist unless there is a pointer.

**Exception - 2**:

You may take address of slice or map only if you're sharing down the callstack to a function that's either named decode or unmarshal.
 
 ### Remember !
 * You do not have a right to make a copy of a value that a pointer points to when it's been shared with you.Assume that it is dangerous to make copies if something has been shared.
 
 * This is critical !
    * Define the data
    * Define the type
    * Get the semantic
    * Write code
    
 * 
    * Choose a semantic
    * Be consistent
    * If it is the wrong one ,np, we'll refactor later.   