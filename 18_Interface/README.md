# Interfaces
* Interfaces are types that declare behaviour.

* Embedding interface ->
```
type ReadWriter interface {
    Reader
    Writer
}
```

* A type *SATISFIES* an interface if it possesses all the methods the interface requires.
* Conceptually, a value of an interface type, or INTERFACE VALUE, has two components,
      a CONCRETE TYPE and a
      VALUE OF THAT TYPE.
  These are called the interface's
      DYNAMIC TYPE and
      DYNAMIC VALUE.
  
  For a statically typed language like Go, types are a compile-time concept, so a type is not a value.
  In our conceptual model, a set of values called TYPE DESCRIPTORS provide information about each type,
  such as its name and methods. In an interface value, the type component is represented by the appropriate
  type descriptor.

```
var w io.Writer
/* w 
 * type: nil
 * value: nil
 */
 
W = os.Stdout
/* w 
 * type: *os.File
 * value: the address where a value of type os.File is stored.
 */
 
w = new(bytes.Buffer)
/* w 
 * type: *bytes.Buffer
 * value: the address where a value of type bytes.Buffer is stored.
 */
 
w = nil
/* w 
 * type: nil
 * value: nil
 */
```