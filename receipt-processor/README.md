# Receipt Processor

A Go web service that processes receipts and calculates reward points based on specific rules.

## Features

- Process receipts and calculate points based on various rules
- In-memory storage for receipts
- RESTful API endpoints for receipt processing and points calculation
- Input validation according to OpenAPI specification
- Docker support

## Prerequisites

- Go 1.21 or higher


## Installation

1. Clone the repository:
```bash
git clone https://github.com/sahith0100/fetch_backend_dev.git
cd receipt-processor
```

2. Install dependencies:
```bash
go mod download
```

## Running the Application

### Using Go

```bash
go run main.go
```
The server will start on `http://localhost:8080`

### Using Docker

```bash
docker-compose up --build
```

## API Endpoints

### Process Receipt
- **URL**: `/receipts/process`
- **Method**: `POST`
- **Request Body**: JSON Receipt object
- **Success Response**: JSON containing receipt ID

## Input Validation

The service validates:
- Retailer name format (alphanumeric, spaces, hyphens, and & only)
- Date format (YYYY-MM-DD)
- Time format (HH:mm, 24-hour)
- Price format (must have exactly 2 decimal places)
- Item description format (alphanumeric, spaces, and hyphens only)

## Error Handling

The service returns appropriate HTTP status codes:
- 200: Success
- 400: Invalid input with validation message
- 404: Receipt not found



