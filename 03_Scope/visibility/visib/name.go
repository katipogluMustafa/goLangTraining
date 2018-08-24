package visib

var MyName = "todd"//Starts with uppercase, visible from outside of the package
var yourName = "Future Programmer"//Starts with lowercase, visible from only inside of the package

/* The Function's name starts with uppercase, visible from outside of the package*/
func Name() string {
	return yourName
}