package main

func main() {
	d := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var c int = 2
	for _, v := range d {
		if v%c == 0 {
			println(v, "is even")
		} else {
			println(v, "is odd")
		}
	}
}
