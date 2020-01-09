package main

import "fmt"

func main(){
	numbers:= []int{0,1,2,3,4,5,6,7,8,9,10}
	for _,value := range numbers{
		if value%2==0{
			fmt.Println("Even")
		}else{
			fmt.Println("Odd")
		}
	}
}

func NumberOfOddEven(numbers []int) (int,int) {
	odd,even:= 0,0
	for _,value := range numbers{
		if value%2==0{
			even++
		}else{
			odd++
		}
	}
	fmt.Println(odd,even)
	return odd,even
}