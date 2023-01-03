package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	assert.Len(t, d, 16)
}
