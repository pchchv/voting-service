package main

type Poll struct {
	title   string
	options map[string]int
}

func main() {
	server()
}
