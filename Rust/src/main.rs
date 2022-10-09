use std::time::{SystemTime, UNIX_EPOCH};

//
// Using the Previous Python and Golang examples,
// We can assume the GET, PUT, POST requests are
// The same speeds when their bodies are empty.
//

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