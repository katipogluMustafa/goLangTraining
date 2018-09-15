# Struct

* There is no implicit conversion when we are talking about concrete data(named types).
```
type bill struct {
        flag    bool
        counter int16
        pi      float32
}
type alice struct {
        flag    bool
        counter int16
        pi      float32
}

fun main(){
    var b bill
    var a alice
//  b = a  // This won't work
    b = bill(a) 
    /*
     * You need to explicitly convert a to b to assign a to b.  
     */
}
```

* Anonymous Struct With Zero Values
```
func main(){
    var e1 struct {
        flag    bool
        counter int16
        pi      float32
    }
}
```
* Anonymous Struct With Value Initializing
```
func main(){
    e1 := struct {
        flag    bool
        counter int16
        pi      float32
    }{
        true,
        1,
        3.141592
    }
}
```

* When we are talking about anonymous structs , you can implicitly convert things.

```
package main

import "fmt"

type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type alice struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	e1 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		true,
		1,
		3.141592,
	}

	var b bill
	var a alice

	b = bill(a) // With named structs, you need explicit conversion
	b = e1      // with anonymous structs, implicitly conversion happens

	fmt.Println(b)
}

```
