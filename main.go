package main

type Poll struct {
	title   string
	options map[string]int
}

func creator(title string, options []string) Poll {
	o := make(map[string]int)
	for _, v := range options {
		o[v] = 0
	}
	return Poll{title: title, options: o}
}

func main() {
	server()
}
