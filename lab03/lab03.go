package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: implement a calculator
	data :=strings.Split(r.URL.Path,"/")
	if len(data) !=4{
		fmt.Fprintf(w,"Error!")
		return
	}
	op := data[1]
	number1, error1 := strconv.Atoi(data[2])
    number2, error2 := strconv.Atoi(data[3])
	if error1 != nil || error2 != nil {
        fmt.Fprintf(w, "Error!")
        return
    }
	output :=0
	reminder :=0
	if  op == "add"{
		output = number1+number2 
		str := fmt.Sprintf("%d + %d = %d",number1,number2,output)
		fmt.Fprintf(w, str)
		return
	}else if op == "sub"{
		output = number1-number2
		str := fmt.Sprintf("%d - %d = %d",number1,number2,output)
		fmt.Fprintf(w, str)
		return
	}else if op == "mul"{
		output = number1*number2
		str := fmt.Sprintf("%d * %d = %d",number1,number2,output)
		fmt.Fprintf(w, str)
		return
	}else if op == "div"{
		if number2 == 0{
			fmt.Fprintf(w, "Error!")
			return
		}
		output = number1/number2
		reminder = number1%number2
		str := fmt.Sprintf("%d / %d = %d, reminder = %d",number1,number2,output,reminder)
		fmt.Fprintf(w, str)
		return
	}else{
		fmt.Fprintf(w, "Error!")
		return
	}
	
	
}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
