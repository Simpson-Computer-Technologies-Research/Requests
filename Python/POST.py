import requests, time

# // Start time
start: int = time.time()

# // Send the http request
requests.post("http://httpbin.org/post")

# // Print how long the request took
print(time.time()-start)


# // 
# // Average Result: 64.0461ms
# //
