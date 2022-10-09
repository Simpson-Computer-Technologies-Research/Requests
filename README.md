# Requests
![banner (1)](https://user-images.githubusercontent.com/75189508/194771973-dec86e5e-ac0c-4dea-a8cb-3f48514c2709.png)
Which is faster? Sending http requests in Go, Python or Rust?

# Benchmarks

```rust

// Python
GET:  		Average Result: 64.0871ms
POST: 		Average Result: 64.0461ms
PUT:  		Average Result: 64.4185ms
THREADED: 	Average Result: 6.28s

// Golang
GET:  		Average Result: 57.3495ms
POST: 		Average Result: 59.2759ms
PUT:  		Average Result: 57.793ms
GOROUTINES: 	Average Result: 289.2099ms

// Rust
GET/PUT/POST: 	Average Result: 183ms
THREADED:	Undetermined

```

# Functions

<h3>Python</h3>

```py

requests.get("http://httpbin.org/get")
requests.post("http://httpbin.org/get")
requests.put("http://httpbin.org/get")

# // Send http request function
def main():
    requests.get("http://httpbin.org/get")

# // Create new threads and start them
for _ in range(100):
    threading.Thread(target=main()).start()

```

<h3>Rust</h3>

```rust
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let request = reqwest::get("https://httpbin.org/get");

    // Store the start time of the request
    let start: u128 = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_millis();

    // Send the HTTP GET Request
    request.await?;

    // Store the end time of the request
    let end: u128 = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_millis();

    // Print the time it took to send the http request
    // Average Result: 183ms
    println!("{}ms", end-start);

    // Return Success
    return Ok(())
}
```


<h3>Go</h3>

```go
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
```


# License
MIT License

Copyright (c) 2022 Tristan Simpson

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
