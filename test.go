package main

import (
	"fmt"
)

// func main() {
// 	var dummyData = structs.VerificationData{Name: "David", VerificationLink: "http://localhost:8080/verify"}

// 	utils.SendGomail(structs.Verification, dummyData, "Verification", []string{"davidlou0810@gmail.com", "davidlois0810@gmail.com"}...)
// }

type I interface {
	M()
}

type about struct {
	Name string
	Age  int
}

func (a *about) M() {
	fmt.Println("Testing print")
}

// func update(intSlice about) {
// 	intSlice.Age = 100
// }

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case I:
		fmt.Printf("I")
	case about:
		fmt.Printf("ABOUT")
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func Test() {
	do(21)
	do("hello")
	do(true)

	a := about{"David", 20}

	do(a)

	var i I = &a
	do(i)
}
