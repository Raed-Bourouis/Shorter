package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func hashURL(longURL string) []byte {
	hash := sha256.New()
	hash.Write([]byte(longURL))
	result := hash.Sum(nil)
	return result
}

func base58Encoded(bytes []byte) string {

	encoder := base58.FlickrEncoding
	encoded, err := encoder.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortURL(initialLink string, userId string) string {
	urlHashBytes := hashURL(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
