package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.youtube.com/watch?v=dQw4w9WgXcQ&pp=ygUJcmljayByb2xs"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	// Persist data mapping
	SaveURLMapping(shortURL, initialLink, userUUId)

	// Retrieve initial URL
	retrievedUrl := GetLongURL(shortURL)

	assert.Equal(t, initialLink, retrievedUrl)
}
