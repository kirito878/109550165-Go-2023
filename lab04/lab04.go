package main

import (
	"fmt"
	"html/template"
	// "html/template"
	"log"
	"net/http"
	"strconv"
)

// TODO: Create a struct to hold the data sent to the template

// Define a struct type
type Output struct {
    Expression string
    Result  string
}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: Finish this function
    op := r.URL.Query().Get("op")
    num1 := r.URL.Query().Get("num1")
    num2 := r.URL.Query().Get("num2")
	number1, error1 := strconv.Atoi(num1)
	number2, error2 := strconv.Atoi(num2)
    if error1 != nil || error2 != nil {
        http.ServeFile(w,r,"error.html")
        return
    }
    if op == "add" {
        // Convert the num1 and num2 strings to integers for addition
		result := number1 + number2
		expr := fmt.Sprintf("%d + %d",number1,number2)
		res := fmt.Sprintf("%d",result) 
		var data = Output{
			Expression: expr,
			Result: res,
		}
		template.Must(template.ParseFiles("index.html")).Execute(w,data)
		return
		}else if op == "sub"{
			result := number1-number2
			expr := fmt.Sprintf("%d - %d",number1,number2)
			res := fmt.Sprintf("%d",result)
			var data = Output{
				Expression: expr,
				Result: res,
			}
			template.Must(template.ParseFiles("index.html")).Execute(w,data)
			return
		}else if op == "mul"{
			result := number1*number2
			expr := fmt.Sprintf("%d * %d",number1,number2)
			res := fmt.Sprintf("%d",result)
			var data = Output{
				Expression: expr,
				Result: res,
			}
			template.Must(template.ParseFiles("index.html")).Execute(w,data)
			return
		}else if op == "div"{
			if number2 == 0{
				http.ServeFile(w,r,"error.html")
				return
			}
			result := number1/number2
			expr := fmt.Sprintf("%d / %d",number1,number2)
			res := fmt.Sprintf("%d",result)
			var data = Output{
				Expression: expr,
				Result: res,
			}
			template.Must(template.ParseFiles("index.html")).Execute(w,data)
			return
		}else if op == "gcd"{
			result := gcd(number1,number2)
			expr := fmt.Sprintf("GCD(%d, %d)",number1,number2)
			res := fmt.Sprintf("%d",result)
			var data = Output{
				Expression: expr,
				Result: res,
			}
			template.Must(template.ParseFiles("index.html")).Execute(w,data)
			return
		}else if op == "lcm"{
			result := lcm(number1,number2)
			expr := fmt.Sprintf("LCM(%d, %d)",number1,number2)
			res := fmt.Sprintf("%d",result)
			var data = Output{
				Expression: expr,
				Result: res,
			}
			template.Must(template.ParseFiles("index.html")).Execute(w,data)
			return
		}else{
			http.ServeFile(w,r,"error.html")
			return
		}
}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8084", nil))
}
