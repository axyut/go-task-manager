## URL Shortener

### Description:

A simple URL shortener service in Golang. The service is able to take a long URL as input, generate a short URL, and then redirect users from the short URL to the original long URL.

### Requirements:

- Golang
- Text editor, terminal

### Run:

- Run `go run .` from the root of the project `url_shortner/`.
- Generate short URL:
  - with curl: `curl -L 'http://localhost:8080/shorten' -H 'Content-Type: application/json' -d '{"url":"https://facebook.com"}'`
  - with Postman:
    - Set POST Method
    - Set Url `http://localhost:8080/shorten`
    - Navigate to `Body` section and click `raw` and select `JSON`
    - Type `{"url":"https://facebook.com"}` at code palatte.
    - Send Request.
  - with Postman: Send request to `http://localhost:8080/shorten?url=https://facebook.com` setting Method to POST.
- Use Short URL, by navigating to the link provided at response of `/shorten`.

### Run Tests:

- Run `go test` inside the cmd directory.

### Features:

1. The service exposes two endpoints:
   - POST /shorten: Accepts a JSON payload with a long URL and returns a short URL.
   - GET /{shortCode}: Redirects users to the original long URL based on the short code provided.
1. Uses a simple in-memory data structure to store the mapping between short URLs and long URLs. (assume that the data will fit in memory).
1. The short URL is generated using a base64 encoding of a unique identifier.
1. Implemented proper error handling for cases like invalid URLs, non-existing short codes etc.
1. Unit tests for the packages.
1. A README with instructions on how to run the service and the tests.
1. Uses Go modules for dependency management.

### Includes:

- Code Quality: Clean, readable, and well-organized code.
- Functionality: The service correctly shorten URLs and redirect users to the original long URL.
- Error Handling: Proper handling of edge cases and errors.
- Documentation: A clear README file with instructions on how to run the service and tests.
- Go Best Practices: Adherence to best practices in Golang development.
