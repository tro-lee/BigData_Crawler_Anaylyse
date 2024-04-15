package manager

type PageData struct {
	data  string
	index int
}

type FilmData struct {
	Name    string
	Score   string
	People  string
	Content string
	Comment string
}

type DetailData struct {
	Director string
	Country  []string
	Year     string
	Genre    []string
}

type ResultData struct {
	Name     string
	Score    string
	People   string
	Comment  string
	Director string
	Country  []string
	Year     string
	Genre    []string
}
