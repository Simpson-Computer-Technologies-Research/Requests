import requests, time

# // Start time
start: int = time.time()

# // Send the http request
requests.put("http://httpbin.org/put")

# // Print how long the request took
print(time.time()-start)


# //
# // Average Result: 64.4185ms
# //
