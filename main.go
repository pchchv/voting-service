package main

type Poll struct {
	Title   string         `json:"title"`
	Options map[string]int `json:"options"`
}

func creator(title string, options []string) Poll {
	o := make(map[string]int)
	for _, v := range options {
		o[v] = 0
	}
	return Poll{Title: title, Options: o}
}

func main() {
	server()
}
