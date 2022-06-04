package main

import (
	api "github.com/dhruvit2/url_shortner/api/http"
	radis "github.com/dhruvit2/url_shortner/store/radis"
	shortner "github.com/dhruvit2/url_shortner/shortner"
)
func main() {

	db, err := radis.NewStore()
	if err != nil {
		panic("Failed to get redis connection ", err)
	}

	appl := shortner.New(db)

	r := api.NewRouter(appl)

	go func() {
		err := r.Run(":9808")
		if err != nil {
			panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, unix.SIGINT, unix.SIGTERM)
	<-quit


}