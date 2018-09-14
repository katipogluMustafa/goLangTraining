# Slices

* When we initialize a slice with zero value, we get nil slice.
    * Nil slice has no pointer that points to the underlying array.
    * pointer: nil length: 0 capacity: 0
```
var data []string
```
* Empty Struct literal will give you an empty slice
    * pointer: points to the empty struct length: 0 capactiy: 0
```
data := []string{}
```


* If you slice a slice and make changes on the new slice, the original slice will be changed too.

```
s1 :=  []int{1,2,3,4,5,6}
fmt.Println("s1:", s1) // s1: [1 2 3 4 5 6]
s2 := s1[2:4]
s2[0] = 10
fmt.Println("s1:", s1) // s1: [1 2 10 4 5 6]
```

* Subslice of a slice will point to the same slice from the point it sliced to the end of the original one. 
```
a := []int{1, 2, 3, 4, 5, 6}
b := a[2:4]
c := a[:3]
d := a[3:]
fmt.Println("a: ", a) // a: [1 2 3 4 5 6]
fmt.Println("b: ", b) // b: [3 4 ]
fmt.Println("c: ", c) // c: [1 2 3]
fmt.Println("d: ", d) // d: [4 5 6]
fmt.Println("Capacity of b: ", cap(b)) // Capacity of b: 4
fmt.Println("What b actually sees: ", b[:len(b)]) // What b actually sees: [3 4]
fmt.Println("Actually b: ", b[:cap(b)])// [3 4 5 6]
```    

* Copy() function is used to copy the contents of one slice to another.Copy() will create separate underlying array that is different the the original slice's underlying array."
```
s1 := []int{1, 2, 3, 4, 5, 6}
fmt.Println("s1:", s1) // [1 2 3 4 5 6]
s2 := make([]int, 2)
n := copy(s2, s1[2:4])
fmt.Println("Number of items copied: ", n) // 2
s2[0] = 10
fmt.Println("s1: ", s1) // [1 2 3 4 5 6]
fmt.Println("s2: ", s2) // [10 4]
fmt.Println(s2[:cap(s2)]) // [10 4]
```

* Delete Example
```
a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
a = append(a[:2], a[4:]...)
fmt.Println(a) // [1, 2, 5, 6, 7, 8, 9, 10]
```

#### Delete
```go
a = append(a[:i], a[i+1:]...)
// or
a = a[:i+copy(a[i:], a[i+1:])]
```

#### Delete without preserving order
```go
a[i] = a[len(a)-1] 
a = a[:len(a)-1]

```
**NOTE** If the type of the element is a _pointer_ or a struct with pointer fields, which need to be garbage collected, the above implementations of ` Cut ` and ` Delete ` have a potential _memory leak_ problem: some elements with values are still referenced by slice ` a ` and thus can not be collected. The following code can fix this problem:
> **Cut**
```go
copy(a[i:], a[j:])
for k, n := len(a)-j+i, len(a); k < n; k++ {
	a[k] = nil // or the zero value of T
}
a = a[:len(a)-j+i]
```

> **Delete**
```go
copy(a[i:], a[i+1:])
a[len(a)-1] = nil // or the zero value of T
a = a[:len(a)-1]
```

> **Delete without preserving order**
```go
a[i] = a[len(a)-1]
a[len(a)-1] = nil
a = a[:len(a)-1]
```