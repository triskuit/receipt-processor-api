# Receipt Processor

An example RESTful web service that processes store receipts and calculates reward points based on specific rules. Built with Go and Gin web framework for Fetch.

## Features

- Process receipts and generate unique IDs
- Calculate points based on various receipt attributes
- In-memory storage for receipt points
- Input validation for all receipt fields
- RESTful API endpoints
- Test coverage for business logic

## Prerequisites

- Go 1.19 or higher
- Git (for cloning the repository)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/receipt-processor.git
cd receipt-processor
```

2. Install dependencies:
```bash
go mod download
```

## Usage

1. Start the server:
```bash
go run main.go
```

The server will start on `localhost:8080`.

## Limitations

- Data is stored in memory and will be lost when the service restarts
- No authentication or authorization
- No rate limiting
- Single-instance only (no distributed storage)
