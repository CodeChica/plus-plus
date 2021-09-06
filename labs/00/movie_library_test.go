package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovieLibrary(t *testing.T) {
	aBugsLife := NewMovie("A Bug's Life", Pixar, 1998)
	cars := NewMovie("Cars", Pixar, 2009)
	fantasia := NewMovie("Fantasia", Disney, 1940)
	monstersInc := NewMovie("Monsters Inc.", Pixar, 2001)
	pinocchio := NewMovie("Pinocchio", Disney, 1940)
	toyStory := NewMovie("Toy Story", Pixar, 1995)
	up := NewMovie("Up", Pixar, 2006)

	t.Run("Add", func(t *testing.T) {
		t.Skip()

		library := NewMovieLibrary()
		library.Add(toyStory)
		library.Add(aBugsLife)

		assert.Equal(t, 2, library.TotalCount())
		assert.Equal(t, toyStory, library.All()[0])
		assert.Equal(t, aBugsLife, library.All()[1])
	})

	t.Run("Remove", func(t *testing.T) {
		t.Skip()

		library := NewMovieLibrary()
		library.Add(aBugsLife)
		library.Add(cars)
		library.Add(fantasia)
		library.Add(monstersInc)
		library.Add(pinocchio)
		library.Add(toyStory)
		library.Add(up)

		library.Remove(cars)

		assert.Equal(t, 6, library.TotalCount())

		library.Each(func(movie Movie) {
			assert.NotEqual(t, cars, movie)
		})
	})

	t.Run("FindBy", func(t *testing.T) {
		t.Skip()

		library := NewMovieLibrary()
		library.Add(aBugsLife)
		library.Add(cars)
		library.Add(fantasia)
		library.Add(monstersInc)
		library.Add(pinocchio)
		library.Add(toyStory)
		library.Add(up)

		result := library.FindBy(func(movie Movie) bool {
			return movie.Title == "Cars"
		})

		assert.Equal(t, cars, result)
	})

	t.Run("SortBy", func(t *testing.T) {
		t.Skip()

		library := NewMovieLibrary()
		library.Add(aBugsLife)
		library.Add(cars)
		library.Add(fantasia)
		library.Add(monstersInc)
		library.Add(pinocchio)
		library.Add(toyStory)
		library.Add(up)

		result := library.SortBy(func(x Movie, y Movie) int {
			return strings.Compare(x.Title, y.Title)
		})

		assert.NotEqual(t, nil, result)
		assert.Equal(t, cars, result)
	})
}
