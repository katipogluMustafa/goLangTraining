## Semantic Usage
    Remember, semantic consistancy is everything!

    Value semantics is always first choice, Pointer semantics usage should be exception.  
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
 