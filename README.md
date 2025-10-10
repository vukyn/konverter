# Konverter - MessagePack API Server

A Go-based API server using Fiber framework for encoding and decoding MessagePack data.

## Features

-   ✅ Fiber web framework
-   ✅ CORS enabled
-   ✅ Rate limiting (100 requests per minute per IP)
-   ✅ Health check endpoint
-   ✅ MessagePack encode/decode with base64 and raw bytes support

## API Endpoints

### Health Check

```
GET /health
```

Response:

```json
{
	"success": true,
	"data": {
		"status": "healthy",
		"timestamp": "2024-01-01T00:00:00Z",
		"service": "konverter"
	}
}
```

### Encode MessagePack

```
POST /encode
```

Request body:

```json
{
	"type": "base64|bytes",
	"data": "your_data_here"
}
```

Response:

```json
{
	"success": true,
	"data": {
		"encoded": "base64_encoded_msgpack_data",
		"type": "base64"
	}
}
```

### Decode MessagePack

```
POST /decode
```

Request body:

```json
{
	"type": "base64|bytes",
	"data": "msgpack_data_here"
}
```

Response:

```json
{
	"success": true,
	"data": {
		"decoded": "decoded_data",
		"type": "base64"
	}
}
```

## Usage

### Start the server

```bash
go run main.go
```

The server will start on port 3000 by default, or use the `PORT` environment variable to specify a different port.

### Example requests

#### Encode JSON data as MessagePack

```bash
curl -X POST http://localhost:3000/encode \
  -H "Content-Type: application/json" \
  -d '{
    "type": "bytes",
    "data": "{\"name\": \"John\", \"age\": 30}"
  }'
```

#### Decode MessagePack data

```bash
curl -X POST http://localhost:3000/decode \
  -H "Content-Type: application/json" \
  -d '{
    "type": "base64",
    "data": "base64_encoded_msgpack_data"
  }'
```

## Rate Limiting

The API has rate limiting enabled:

-   100 requests per minute per IP address
-   Returns HTTP 429 (Too Many Requests) when limit is exceeded

## CORS

CORS is enabled for all origins with the following methods:

-   GET, POST, PUT, DELETE, OPTIONS

## Dependencies

-   [Fiber v2](https://github.com/gofiber/fiber) - Web framework
-   [MessagePack v5](https://github.com/vmihailenco/msgpack) - MessagePack encoding/decoding
