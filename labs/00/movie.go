package main

type Movie struct {
	Title         string
	Studio        Studio
	YearPublished int
}

func NewMovie(title string, studio Studio, year int) *Movie {
	return &Movie{}
}

func (movie *Movie) CompareTo(other *Movie) int {
	return 0
}
