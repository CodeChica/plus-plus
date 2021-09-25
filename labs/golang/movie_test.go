package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovie(t *testing.T) {
	t.Run("NewMovie", func(t *testing.T) {
		t.Skip()

		toyStory := NewMovie("Toy Story", Pixar, 1995)

		assert.Equal(t, "Toy Story", toyStory.Title)
		assert.Equal(t, 1995, toyStory.YearPublished)
		assert.Equal(t, Pixar, toyStory.Studio)
	})

	t.Run("Equality", func(t *testing.T) {
		t.Skip()

		toyStory := NewMovie("Toy Story", Pixar, 1995)
		aBugsLife := NewMovie("A Bug's Life", Pixar, 1998)

		assert.Equal(t, aBugsLife, aBugsLife)
		assert.Equal(t, toyStory, NewMovie("Toy Story", Pixar, 1995))
		assert.Equal(t, toyStory, toyStory)
		assert.NotEqual(t, toyStory, aBugsLife)
	})

	t.Run("Compare", func(t *testing.T) {
		t.Skip()

		toyStory := NewMovie("Toy Story", Pixar, 1995)
		aBugsLife := NewMovie("A Bug's Life", Pixar, 1998)

		assert.Equal(t, 0, aBugsLife.CompareTo(aBugsLife))
		assert.Equal(t, -1, aBugsLife.CompareTo(toyStory))
		assert.Equal(t, 1, toyStory.CompareTo(aBugsLife))
	})
}
