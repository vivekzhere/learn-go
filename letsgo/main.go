package main

import (
	"fmt"

	"github.wdf.sap.corp/i332859/learn-go/letsgo/structs"
)

func main() {
	fmt.Println("hello world")

	var myObj structs.MyStruct
	myObj = structs.MyStruct{}
	myObj.Age = 20

	myObj2 := structs.Struct1{
		Name: "vivek",
	}

	fmt.Println(myObj2)

	if myObj2.Name != "vivek" {
		fmt.Println("not me")
	} else {
		fmt.Println("me")
	}
}
