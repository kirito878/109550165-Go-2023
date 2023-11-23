package main

import (
	"fmt"
	"math/big"
	// "strconv"
	"syscall/js"
)

func CheckPrime(this js.Value, args []js.Value) interface{} {
	// TODO: Check if the number is prime
	inputValue := js.Global().Get("value").Get("value").String()

	num, success := new(big.Int).SetString(inputValue, 10)
	if !success {
		return "Invalid number"
	}

	isPrime := num.ProbablyPrime(0)
	resultString := "It's not prime"
	if isPrime {
		resultString = "It's prime"
	}
	js.Global().Get("answer").Set("innerText", resultString)

	return nil
}

func registerCallbacks() {
	// TODO: Register the function CheckPrime
	js.Global().Set("CheckPrime", js.FuncOf(CheckPrime))


}

func main() {
	fmt.Println("Golang main function executed")
	registerCallbacks()

	//need block the main thread forever
	select {}
}
