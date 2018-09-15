package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}

/*

 * A struct is a sequence of named elements, called fields, each of which has a name and a type.
 * Field names may be specified explicitly(IdentifierList) or implicitly(AnonymousField).
 * Within a struct, non-blank field names must be unique.

 *A field declared with a type but no explicitly	 field name is an anonymous field, also called an embedded field or an embedding of the type in the struct.
 * An embedded type must be specified as a type name T or as a pointer to a non-interface type name *T, and T itself may not be a pointer type.
 * The unqualified type name acts as the field name.
 */
