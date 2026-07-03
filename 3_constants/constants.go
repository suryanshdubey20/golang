package main

import "fmt"

//const age = 20

func main(){
	//const name = "golang"
	//const age = 20
	
	//fmt.Println(age)

    const (
		port = 2300
		host = "loaclhost"
	)

	fmt.Println(port, host)

}