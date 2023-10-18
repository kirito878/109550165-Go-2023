package main

import "fmt"
import "strconv"
import "strings"
func main() {
	var n int64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	result := Sum(n)
	fmt.Println(result)
}

func Sum(n int64) string {
	// TODO: Finish this function
	var s int64;
	s = 0;
	exp :=[]string{}
	for i :=int64(1)  ;i <= n ;i++ {
			if i % 7 !=0{
				s+=i
				exp = append(exp, strconv.FormatInt(i, 10))
			}
		}
	result :=strings.Join(exp,"+")
	str := fmt.Sprintf("%s=%d",result,s)

	return str

}
