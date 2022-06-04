package store

type DataStore interface {
	/* We want to be able to save the mapping between the originalUrl
	and the generated shortUrl url
	*/
	SaveUrlMapping(shortUrl string, originalUrl string, userId string) error
	/*
	We should be able to retrieve the initial long URL once the short
	is provided. This is when users will be calling the shortlink in the
	url, so what we need to do here is to retrieve the long url and
	think about redirect.
	*/
	RetrieveInitialUrl(shortUrl string) (string, error)
}