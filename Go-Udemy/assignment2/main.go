package main

func main() {
	t := triangle{
		base:   5,
		height: 3,
	}
	s := square{sideLength: 4}
	printArea(t)
	printArea(s)
}
