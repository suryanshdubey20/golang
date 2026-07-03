package maim

import "fmt"

func main() {
	age := 16
	if age >= 18 {
		fmt.Println("preson is aadult")
	} else {
		fmt.Println("person is not adult")
	}


	//age := 10
    if age >= 18 {
		fmt.Println("person is adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is kid")
	}


	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission{
		fmt.Println("yes")
	}

   
	//we can declaare a variable inside if costruct.
	if age := 15; age >= 18 {
		fmt.Println("person is adult", age)
	}else{
        fmt.Println("person is not adult", age)
	}

// go does not have ternary operator , you will have to use normal  if else 

}
