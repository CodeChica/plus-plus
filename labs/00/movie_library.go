package main

type MovieLibrary struct {
}

func NewMovieLibrary() *MovieLibrary {
	return &MovieLibrary{}
}

func (library *MovieLibrary) Add(movie *Movie) {
}

func (library *MovieLibrary) TotalCount() int {
	return 0
}
