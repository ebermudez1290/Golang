package main

import "testing"

func TestOddEven(t *testing.T){
	numbers:= []int{1,15,16,18,49,68,44,25,78,89,42,33,1,51,51}
	odd,even:= NumberOfOddEven(numbers)

	if(odd!=9 || even!=6){
		t.Errorf("We expected 6 odds and 9 evens and received %v %v",odd,even)
	}
}