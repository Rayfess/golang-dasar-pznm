package main

func main() {
	const (
		a = 5
		b = 20
		c = 100
		z = 12
	)

	var d = a - b - c
	var e = -a + b*c
	var f = c / b
	var g = z % a

	println(d)
	println(e)
	println(f)
	println(g)
}