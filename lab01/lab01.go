package main

import "fmt"

func main() {
	fmt.Println("Welcome to Simple Calculator")

	var a, b int64
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("Add:", Add(a, b))
	fmt.Println("Subtract:", Sub(a, b))
	fmt.Println("Multiply:", Mul(a, b))
	fmt.Println("Divide:", Div(a, b))
}

// TODO: Create `Add`, `Sub`, `Mul`, `Div` function here
func Add(a,b int64)(int64){
	var c int64
	c = a+b
	return c
}
func Sub(a,b int64)(int64){
	var c int64
	c = a-b
	return c
}
func Mul(a,b int64)(int64){
	var c int64
	c = a*b
	return c
}
func Div(a,b int64)(int64){
	var c int64
	c = a/b
	return c
}