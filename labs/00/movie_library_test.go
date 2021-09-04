package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovieLibrary(t *testing.T) {
	toyStory := NewMovie("Toy Story", Pixar, 1995)
	aBugsLife := NewMovie("A Bug's Life", Pixar, 1998)
	up := NewMovie("Up", Pixar, 2006)
	cars := NewMovie("Cars", Pixar, 2009)
	monstersInc := NewMovie("Monsters Inc.", Pixar, 2001)

	fantasia := NewMovie("Fantasia", Disney, 1940)
	pinocchio := NewMovie("Pinocchio", Disney, 1940)

	t.Run("Add", func(t *testing.T) {
		t.Skip()

		library := NewMovieLibrary()
		library.Add(toyStory)
		library.Add(aBugsLife)

		assert.Equal(t, 2, library.TotalCount())
	})

}
