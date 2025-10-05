package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "12345678-abcd-4321-efgh-87654321abcd"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://golang.org/doc/effective_go"
	shortLink_1 := GenerateShortURL(initialLink_1, UserId)

	initialLink_2 := "https://pkg.go.dev/testing"
	shortLink_2 := GenerateShortURL(initialLink_2, UserId)

	initialLink_3 := "https://github.com/stretchr/testify"
	shortLink_3 := GenerateShortURL(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "e2vUGLzN")
	assert.Equal(t, shortLink_2, "kRGF9MYi")
	assert.Equal(t, shortLink_3, "cVeoHdzM")
}
