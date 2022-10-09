import requests, time

# // Start time
start: int = time.time()

# // Send the http request
requests.get("http://httpbin.org/get")

# // Print how long the request took
print(time.time()-start)


# //
# // Average Result: 64.0871ms
# //
