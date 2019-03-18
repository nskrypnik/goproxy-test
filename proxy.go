package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"time"
)

/*
	Utilities
*/

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

/*
	Getters
*/

// Get the port to listen on
func getListenAddress() string {
	port := getEnv("PORT", "1338")
	return ":" + port
}

/*
	Reverse Proxy Logic
*/

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	proxyUrl := getEnv("PROXY_TO", "http://localhost:1330")
	extraHeader := getEnv("EXTRA_HEADER", "PROXY_1")
	t := time.Now().UnixNano()
	url, _ := url.Parse(proxyUrl)
	req.Header.Set(extraHeader, strconv.FormatInt(t, 10))
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(res, req)
}

/*
	Entry
*/

func main() {

	// start server
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}
