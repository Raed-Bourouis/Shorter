# Shorter

Shorter is a simple and efficient URL shortener built with Go. This project serves as a hands-on learning experience to explore the fundamentals of the Go programming language and its ecosystem.

## Features

- **URL Shortening**: Quickly generate unique short links for long URLs.
- **Redis Integration**: Persistent and fast storage for URL mappings.
- **Base58 Encoding**: Elegant and compact short URLs.
- **Lightweight and Scalable**: Built with the Gin web framework for performance.

## Getting Started

### Prerequisites

- Go 1.25 or higher
- Redis server

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Raed-Bourouis/Shorter.git
   cd Shorter
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Start the Redis server:
   ```bash
   redis-server --port 5000
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

   The server will be available at `http://localhost:9000`.

## Usage

### Generate a Short URL

You can interact with the server to generate and retrieve short URLs. Example request:

```bash
curl --request POST \
--data '{
 
    "long_url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ&pp=ygUJcmljayByb2xs",
    "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"

}' \
  http://localhost:9000/create-short-url
```

Response : 
```json
{
    "message": "short url created successfully",
    "short_url": "http://localhost:9000/2n2XxSvU"
}

```


## Project Structure

- `main.go`: Entry point for the application, sets up the web server.
- `store/`: Handles Redis-based storage operations for URLs.
- `shortener/`: Contains the logic for generating short URLs.

## Why This Project?

This repository was created to:
- Gain hands-on experience with Go.
- Learn how to build a web application using the Gin framework.
- Explore Redis as a data storage solution in Go.


## Acknowledgement
During this project, I took learnings from the tutorial provided by [eddywm.com](eddywm.com)
