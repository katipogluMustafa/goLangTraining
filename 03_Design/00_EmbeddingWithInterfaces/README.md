# Mechanics Of Interfaces

* Interfaces give us the ability to have polymorphism.

* Polymorphism means that a piece of code changes its behaviour depending on the concrete data that's operating on.

* When should a piece of data have behaviour ? 
    * When we need to implement polymorphism and that polymorphism is gonna give us that levels of decoupling.
* Interfaces are not real.They only define a method set of behaviour.They define contract of behaviour.
* Interfaces are about behaviour.You should not have interfaces that describe thins.
* Interface types are valueless.
---
* If you are working with a value of type T,**only those methods using value receivers belong to the method set for this value**.
* If you are working with pointers, your pointer receiver methods and your value receiver methods, all the methods you declare for that concrete type they exist for pointers.

* You get full method declarations for your pointer data
* You get only your value methods for your value data.In other words, these pointer receiver methods are left out of the methods set for values.

[See Example 1](Example1.go)

---
* If the pointer semantics were the choice ,then you can only share.
* If you choose value semantics, then you can safely make copies.Also there are also times sharing is safe too.
* It's never safe to make a copy of a value pointer points to.
---