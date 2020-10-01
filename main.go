package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"io/ioutil"
)


func getEnv(key, fallback string) string {
	value, foundValue := os.LookupEnv(key)
	if foundValue {
		return value
	}
	return fallback
}

func getNewGifURL() string {
	// make a sample HTTP GET request
	res, err := http.Get("http://giphy-proxy.solutions.656.mba" )

	// check for response error
	if err != nil {
		log.Fatal( err )
	}

	// read all response body
	data, _ := ioutil.ReadAll( res.Body )

	// close response body
	res.Body.Close()
	return string(data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	myGifURL := getNewGifURL()
	myNotFunHTML := `<html>
<style>body{background-color: coral;}</style>	
	
<h1>This is my rad OTTER page!</h1>
<div>
<img src="` + myGifURL + `">
<form action="">
<button type="submit">display gif</button>
</form>
</div>
</html>
`
	fmt.Fprintf(w, myNotFunHTML)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+getEnv("PORT", "8080"), nil)
}
