package utils

import (
	"regexp"
	"strings"
)

func ProcessContent(text string) *DetailData {
	directorRegex := regexp.MustCompile(`导演:\s*([^ ]+)`)
	directorMatches := directorRegex.FindStringSubmatch(text)
	director := directorMatches[1]

	countryRegex := regexp.MustCompile(`/&nbsp;(.*)&nbsp;/`)
	countryMatches := countryRegex.FindStringSubmatch(text)
	country := process(countryMatches[1])

	genreRegex := regexp.MustCompile(`;([^;]*)$`)
	genreMatches := genreRegex.FindStringSubmatch(text)
	genre := process(genreMatches[1])

	yearRegex := regexp.MustCompile(`([0-9]{4})`)
	yearMatches := yearRegex.FindStringSubmatch(text)
	year := yearMatches[1]

	return &DetailData{
		Director: director,
		Country:  country,
		Year:     year,
		Genre:    genre,
	}
}

func process(text string) []string {
	result := strings.Fields(text)
	return result
}
