package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovieLibrary(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		t.Skip()

		toyStory := NewMovie("Toy Story", Pixar, 1995)
		aBugsLife := NewMovie("A Bug's Life", Pixar, 1998)

		library := NewMovieLibrary()
		library.Add(toyStory)
		library.Add(aBugsLife)

		assert.Equal(t, 2, library.TotalCount())
	})
}
