package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Main Function
func main() {
	var client *http.Client = &http.Client{}

	GET(client)  // Average Result: 57.3495ms
	POST(client) // Average Result: 59.2759ms
	PUT(client)  // Average Result: 57.793ms

	// Sending 100 HTTP GET Requests
	// Average Result: 289.2099ms
	GET_USING_GOROUTINES(client)
}

// Sending an HTTP GET Request
func GET(client *http.Client) {
	// Check out https://github.com/Simpson-Computer-Technologies-Research/DeclarationSpeeds
	// To see why you should use 'var ... type' instead of ':='
	var (
		startTime time.Time = time.Now()
		req, _              = http.NewRequest("GET", "http://httpbin.org/get", bytes.NewBuffer([]byte{}))
	)
	client.Do(req)
	fmt.Println(time.Since(startTime))
}

// Sending an HTTP POST Request
func POST(client *http.Client) {
	var (
		startTime time.Time = time.Now()
		req, _              = http.NewRequest("POST", "http://httpbin.org/post", bytes.NewBuffer([]byte{}))
	)
	client.Do(req)
	fmt.Println(time.Since(startTime))
}

// Sending an HTTP PUT Request
func PUT(client *http.Client) {
	// set the HTTP method, url, and request body
	var (
		startTime time.Time = time.Now()
		req, _              = http.NewRequest("PUT", "http://httpbin.org/put", bytes.NewBuffer([]byte{}))
	)
	client.Do(req)
	fmt.Println(time.Since(startTime))
}

// Sending an HTTP GET Request
func GET_USING_GOROUTINES(client *http.Client) {
	// Check out https://github.com/Simpson-Computer-Technologies-Research/DeclarationSpeeds
	// To see why you should use 'var ... type' instead of ':='
	var (
		startTime time.Time       = time.Now()
		wg        *sync.WaitGroup = &sync.WaitGroup{}
	)

	// Starting 100 Goroutines
	for i := 0; i < 100; i++ {
		wg.Add(1)

		// The goroutine
		go func() {
			defer wg.Done()
			var req, _ = http.NewRequest("GET", "http://httpbin.org/get", bytes.NewBuffer([]byte{}))
			client.Do(req)
		}()
	}

	// Wait for all goroutines to finish before
	// Printing the time it took
	wg.Wait()
	fmt.Println(time.Since(startTime))
}
