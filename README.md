# url-shortener
`url-shortener` a backend app which give shortened url for a full url.

### Prerequisite
- golang (`go1.20` or above)
- redis (`6.0`)
- make (optional)
- Docker and docker-compose (optional)

### Installation
- Clone repo
    `git clone https://github.com/snghnaveen/url-shortener.git`

Using docker-compose (__recommended__)
- Run docker-compose file
    `docker-compose up --build`

Using go command
- Make sure redis is running OR set `ENVIRONMENT` variable to `testing` to skip redis server (It will use external package to mimic redis calls)
- Verify and if required update the configs in [application.env](./application.env)
- Install packages
    `go mod tidy`
- Run app
    `go run main.go`

### Usage
Following are the APIs and their example response.

1. Shorten given URL
    Curl : 
    ```bash
    curl --location 'http://localhost:8080/v1/api/shorten' \
    --header 'Content-Type: application/json' \
    --data '{
    "url": "https://snghnaveen.github.io"
    }'
    ```
    Response : 
    ```json
    {
        "error": false,
        "data": {
            "shorten_key": "qNGbdjC7",
            "shorten_url": "localhost:8080/resolve/qNGbdjC7"
        }
    }
    ```

1. Get URL from shorten key
    Curl :
    ```bash
    curl --location 'http://localhost:8080/v1/api/resolve/qNGbdjC7?type=json'
    ```
    Response :
    ```json
    {
    "error": false,
    "data": {
        "url": "https://snghnaveen.github.io"
    }
    }
    ```
1. Get metrics
    Curl :
    ```bash
    curl --location 'http://localhost:8080/v1/api/metrics-top-requested'
    ```
