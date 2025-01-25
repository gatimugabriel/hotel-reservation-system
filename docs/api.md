# API Documentation

## Base URL
```
https://api.hotelreservation.gabu/v1
```

## Authentication
API endpoints that require auth use an issued accessToken(JWT) during signin. These endpoints are marked to help you know they require auth

```
Authorization: Bearer <your_accessToken>
```

## API Endpoints

### Authentication

#### Register User
```http
POST /auth/register
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "securepassword",
    "firstName": "John",
    "lastName": "Doe",
    "phone": "+1234567890"
}
```

Response:
```json
{
    "status": "success",
    "data": {
        "userId": "uuid-here",
        "email": "user@example.com",
        "firstName": "John",
        "lastName": "Doe"
    }
}
```

#### Login
```http
POST /auth/login
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "securepassword"
}
```

Response:
```json
{
    "status": "success",
    "data": {
        "token": "jwt-token-here",
        "expiresIn": 3600
    }
}
```

### Room Management

#### Get Available Rooms
```http
GET /rooms/available?checkIn=2024-01-01&checkOut=2024-01-05
```

Response:
```json
{
    "status": "success",
    "data": {
        "rooms": [
            {
                "id": "room-uuid",
                "type": "DOUBLE",
                "number": "201",
                "price": 150.00,
                "amenities": ["TV", "AC", "WiFi"],
                "maxOccupancy": 2
            }
        ]
    }
}
```

#### Create Booking
```http
POST /bookings
Content-Type: application/json

{
    "roomId": "room-uuid",
    "checkIn": "2024-01-01",
    "checkOut": "2024-01-05",
    "guests": [
        {
            "firstName": "John",
            "lastName": "Doe",
            "email": "john@example.com"
        }
    ],
    "paymentDetails": {
        "cardNumber": "4111111111111111",
        "expiryMonth": "12",
        "expiryYear": "2025",
        "cvv": "123"
    }
}
```

Response:
```json
{
    "status": "success",
    "data": {
        "bookingId": "booking-uuid",
        "totalAmount": 600.00,
        "status": "CONFIRMED",
        "checkIn": "2024-01-01",
        "checkOut": "2024-01-05"
    }
}
```

### Error Responses

#### 400 Bad Request
```json
{
    "status": "error",
    "error": {
        "code": "INVALID_INPUT",
        "message": "Invalid date format for check-in"
    }
}
```

#### 401 Unauthorized
```json
{
    "status": "error",
    "error": {
        "code": "UNAUTHORIZED",
        "message": "Invalid or expired token"
    }
}
```

#### 404 Not Found
```json
{
    "status": "error",
    "error": {
        "code": "NOT_FOUND",
        "message": "Room not found"
    }
}
```

## Rate Limiting
API requests are limited to 100 requests per minute per IP address.

## Pagination
For endpoints that return lists, pagination is supported using the following query parameters:
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 10, max: 100)

Example:
```http
GET /reservations?page=2&limit=20
```