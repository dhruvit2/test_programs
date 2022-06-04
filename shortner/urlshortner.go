package shortner

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

type App interface {
	// define methods
	GenerateShortLink(initialLink string, userId string) (string, error)
	RetrieveInitialUrl(shortUrl string) (string, err)
}

type app struct {
	store         store.DataStore
}

// NewApp initialize a new notification_mgr App
func New(ds store.DataStore) (App, error) {

	app := &app{
		store:         ds,
	}
	return app, nil
}

func (a *app) GenerateShortLink(initialLink string, userId string) (string, error) {
	urlHashBytes := sha256Of(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString, err := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	if err != nil {
		return "", err
	}

	err = a.store.SaveUrlMapping(finalString[:8], creationRequest.LongUrl, creationRequest.UserId)
	if err != nil {
		return "", err
	}

	return finalString[:8], nil
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (a *app) RetrieveInitialUrl(shortUrl string) (string, err) {
	return a.store.RetrieveInitialUrl(shortUrl)
}