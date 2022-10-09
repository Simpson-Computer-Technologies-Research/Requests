import threading, requests, time

# //
# // Send 100 http requests using multithreading
# //

# // Start time
start: int = time.time()

# // Function to Send the http request
def main():
    requests.get("http://httpbin.org/get")

# // Create new threads and start them
for _ in range(100):
    threading.Thread(target=main()).start()

# // Print how long the request took
print(time.time()-start)

# //
# // Average Result: 6.28s
# //




