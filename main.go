package main

import (
	"fmt"
	"net/http"
	"os"
)


func getEnv(key, fallback string) string {
	value, foundValue := os.LookupEnv(key)
	if foundValue {
		return value
	}
	return fallback
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+getEnv("PORT", "8080"), nil)
}
