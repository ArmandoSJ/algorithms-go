package main

import(
	"fmt"
	"./algorithmfiles"
)

func main(){
	lstOfNumbers := []int{2, 3, 5, 7, 11, 13}
	vTarget := 10
	output := two_number_sum.sumTwoNumbers(listOfNumbers, vTarget )
	fmt.Println(output)
	fmt.Println(lstOfNumbers, vTarget)
}