## Grouping Types

* Go is about convention over configuration.
* Go says configuration is limiting.
* Who is about not who we are,but go is about what we do.
* If we group things together by what we do,not by who we are, then we have a level of diversity.
* Go is very much about let's not limit ourselves on configuration but let's focus what we do.
---
* We only want types where we need values of that type.

 Notes:

 * Declare types that represent something new or unique.
 * Validate that a value of any type is created or used on its own.
 * Embed types to reuse existing behaviors you need to satisfy.
 * Question types that are an alias or abstraction for an existing type.
 * Question types whose sole purpose is to share common state.
 * let's keep concrete types as persons, places, things.
 
 Example: [Grouping Types](01_GroupingTypes.go)
 
 ---
 
 ```go
type handle int

func foo(h handle)
```

* When I see something like this, I ask myself, is handle something new or unique or you just trying to abstract concept that handle is integer.
    * If I don't see a method set after this type, its a good indication that maybe we are trying to do some sort of aliasing.
    * remember constants of a kind can be implicitly converted for example when we send 10 as a argument to foo function(foo(10)), it would be compilable.Since you can't get protection to just accept handle type, especially type like handle which is based on built-in type then I don't want you do it at all.If you can't have 100% of the time, you can't have it!
    * For me better API design here for handle, variable name called handle and type being int because handles are really are int.Don't abstract or alias what the data actually is.
    ```go
    func(handle int)    
    ```
    * So if you can't define a type truly represents something new, somethings that isn't truly base type, then I don't want you to have it.

```go
// Duration form time package
type Duration int64
```

* Duration doesn't represent a integer, it represents nanoseconds of time.As you can see the difference between handle and duration, duration does represent something new and there is actually method sets around that.

---    