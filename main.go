package main

type Poll struct {
	title   string
	options map[string]int
}

func creator(title string, options []string) Poll {
	var p Poll
	return p
}

func main() {
	server()
}
