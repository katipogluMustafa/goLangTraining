package main

import ( /* Importing just in the file level, can not be seen even from diff file in the same package*/
	"fmt"
	"goLangTraining/03_Scope/visibility/visib"
)

func main(){
	fmt.Println(visib.MyName)
	visib.PrintName()
}
