# Unitag : QR code resolver API

## Overview

This is a Go-based URL redirection service built using the Echo framework. The service accepts a 6-character identifier and redirects the user based on:

- Language preferences (English, French, Spanish)

- Browser/User-Agent detection (Chrome/Android -> Google, Safari/iOS -> Apple)

## Features

- URL redirection based on custom identifier.

- Supports language-based redirection.

- Written in Golang using Echo framework for high performance.

- Unit tests included.

## Prerequisites

Ensure you have the following installed before proceeding:

- Go

- Git

## Installation & Setup

1️. Clone the Repository

```bash
git clone https://github.com/rohit-shetty-94/unitag.git
cd unitag
```

2️. Install Dependencies

```bash
go mod tidy
```

3. Run the Application

```bash
go run main.go
```

The application will start on port 3200 by default.

4️. Test the Redirection

Try accessing the service using:

```bash
http://localhost:3200/2bXk5q
```

It should redirect you based on the predefined rules.


## API Endpoints

1️. Redirect Request

GET /:id

Description: Accepts a 6-character ID and redirects accordingly.

Example Request:

```bash
GET http://localhost:3200/2bXk5q
```

QR Id's
```bash
	2bXk5q : Unitag
	ax3Goo : Google
	ax4App : Apple
```

Example Responses:

- Valid ID & English Language (Chrome) → Redirects to https://www.google.com/?hl=en

- Valid ID & French Language (Safari) → Redirects to https://www.apple.com/fr

- Invalid ID (length not 6) → Returns 400 Bad Request with error JSON.

## Testing with Postman

When testing the API in Postman, you need to set the following headers:

`Accept-Language`: The language you want to test (e.g., en, fr, es).

`User-Agent`: Set this to the user-agent of the browser or device you want to test (e.g., Chrome, Safari, Android).

Example Postman Request:
URL: http://localhost:3200/2bXk5q

Headers:

Accept-Language: en (or fr, es)

User-Agent: Chrome (or android, safari, iphone, ipad)

## Running Unit Tests

To run the test suite:

```bash
go test ./...
```

This will execute all test cases from the project.

## Author

**[Rohit Shetty]**\
🚀 **GitHub**: [rohit-shetty-94](https://github.com/rohit-shetty-94)\
💎 **Email**: [shettyrohit61@gamil.com](mailto\:shettyrohit61@gamil.com)


postman steps

add google/apple/unitag code in readme