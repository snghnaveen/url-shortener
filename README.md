# url-shortener
`url-shortener` a backend app which give shortened url for a full url.

### Prerequisite
- golang (`go1.20` or above)
- redis (`6.0`)
- make (optional)
- Docker and docker-compose (optional)

### Installation
Clone repo
    ```
    git clone https://github.com/snghnaveen/url-shortener.git
    ```

---
Using docker-compose (__recommended__)
- Run docker-compose file
    ```
    docker-compose up --build
    ```
---
Using go command
- Make sure redis is running OR set `ENVIRONMENT` variable to `testing` to skip redis server (It will use external package to mimic redis calls)
- Verify and if required update the configs in [application.env](./application.env)

- Run app 
    - Run make file
    ```
    make run
    ```

    OR 

    - Install packages
    ```
    go mod tidy
    ```
    - Run app
    ```
    go run main.go
    ```

### Running test
- Run tests
    ```
    ENVIRONMENT=testing go test ./...
    ```
    or
    ```
    make test
    ```

You can github workflow the test case results.

### Project Structure

```bash

├── Dockerfile.dev <--- Docker file for dev, uses air to run go app
├── Makefile
├── README.md
├── application.env <--- app config
├── db <--- Redis connection
│   ├── 
│   ├── 
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go <--- main file
├── pkg
│   ├── resolver <--- service level logic
│   │   ├── 
│   │   └──
│   ├── rest <--- REST api related common code
│   │   ├──
│   │   └──
├── routers
│   ├── api
│   │   └── v1 <--- API group
│   ├── routes.go <--- Registered endpoints
└── util <--- common code (logger, config etc) 
    ├── 
    ├── 

```
### Usage
Following are the APIs and their example response.

- Shorten given URL
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
---
- Get URL from shorten key
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
    Note : Remove query param `type` to see  actual redirection.
---
- Get metrics
    ```bash
    curl --location 'http://localhost:8080/v1/api/metrics-top-requested'
    ```
    Response :
    ```json
    {
        "error": false,
        "data": [
            {
                "rank": 1,
                "score": 100,
                "domain": "https://snghnaveen.1.io"
            },
            {
                "rank": 2,
                "score": 50,
                "domain": "https://snghnaveen.2.io"
            },
            {
                "rank": 3,
                "score": 33,
                "domain": "https://snghnaveen.3.io"
            }
        ]
    }
    ```
