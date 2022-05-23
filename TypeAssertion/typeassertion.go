package main

import "fmt"

func main() {
	var interfaceVariable interface{}
	interfaceVariable = 1
	fmt.Println(interfaceVariable)

	intVariable := 0
	fmt.Println(intVariable)

	//get data from an interface variable and check type convertion status
	intVariable, isSuccecc := interfaceVariable.(int)
	fmt.Println(intVariable, isSuccecc)

	//get variable type
	//this line have error : fmt.Println( interfaceVariable.(type))
	//because we can use .(type) just in switch

	//this line have error : switch intVariable.(type){}
	//because intVariable (variable of type int) is not an interfacecompilerInvalidTypeSwitch

	switch interfaceVariable.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}
