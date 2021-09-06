package main

type MovieLibrary struct{}

func NewMovieLibrary() *MovieLibrary {
	return &MovieLibrary{}
}

func (library *MovieLibrary) Add(movie *Movie) {
}

func (library *MovieLibrary) Remove(movie *Movie) {
}

func (library *MovieLibrary) TotalCount() int {
	return 0
}

func (library *MovieLibrary) All() []Movie {
	return []Movie{}
}

func (library *MovieLibrary) Each(visitor func(Movie)) {
}

func (library *MovieLibrary) FindBy(predicate func(Movie) bool) *Movie {
	return nil
}

func (library *MovieLibrary) SortBy(comparer func(Movie, Movie) int) []Movie {
	return []Movie{}
}
