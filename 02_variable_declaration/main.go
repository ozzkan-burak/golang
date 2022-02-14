package main

import "fmt"

func main() {

	// var name string // var - name of variable - static type

	// name = "Burak"

	//var age int = 45

	//fmt.Println(name, age)
	//fmt.Println(age)

	var isHaveJob, isMarried bool = true, true
	var name, lastName string = "Burak", "Özkan"

	// değişken tanımlamanın kısayolu
	job := "Software Engineer"

	fmt.Println(isHaveJob, isMarried)
	fmt.Println(name, lastName)
	fmt.Println(job)

}
