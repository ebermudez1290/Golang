package main

import "fmt"

func main(){
	colors:= map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	var myColors map[string]string


	numbers:= make(map[int]string)
	numbers[1] = "one"
	delete(numbers,1)

	printMap(colors)
	fmt.Println(myColors)
	fmt.Println(numbers)
}

func printMap(c map[string]string){
	for key,value:= range c{
		fmt.Println(key,value)
	}
}