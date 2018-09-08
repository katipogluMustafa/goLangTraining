# Slices

* If you slice a slice and make changes on the new slice, the original slice will be changed too.

```
s1 :=  []int{1,2,3,4,5,6}
fmt.Println("s1:", s1) // s1: [1 2 3 4 5 6]
s2 := s1[2:4]
s2[0] = 10
fmt.Println("s1:", s1) // s1: [1 2 10 4 5 6]
```

* 
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