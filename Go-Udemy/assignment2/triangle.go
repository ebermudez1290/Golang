package main

type base float64
type triangle struct {
	height float64
	base
}

func (t triangle) getArea() float64 {
	return 0.5 * float64(t.base) * t.height
}
