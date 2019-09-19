package function

func Add() func(int) int {
	var x int

	return func(diff int) int {
		x += diff
		return x
	}
}
