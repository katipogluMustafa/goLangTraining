package visib

var MyName = "todd"//Starts with uppercase, visible from outside of the package
var yourName = "Future Programmer"//Starts with lowercase, visible from only inside of the package

/* The Function's name starts with lowercase, visible only side of the same package*/
func name() string {
	return yourName
}