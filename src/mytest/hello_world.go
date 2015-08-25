package main

var a = "G"

func main() {
	// println("Hello", "World")
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a = "O"
	print(a)
}
