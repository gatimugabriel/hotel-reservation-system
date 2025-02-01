# API Documentation

You can view and interact with the api here. 
```
https://elements.getpostman.com/redirect?entityId=24249138-a653a984-406d-47d2-b479-dd0b192444d7&entityType=collection
```

Once accessed, just ensure sure you set the `base_url` variable to:

``https://localhost:8080/api/v1`` - to interact with LOCAL application that you setup earlier

``https://hrs-nexus.onrender.com/api/v1`` - to interact with hosted application

## Authentication
Many endpoints require authentication using JWT tokens. To authenticate:

1. Register a new user account
2. Login with your credentials
3. Include the JWT token in subsequent requests using the Authorization header:
   `Authorization: Bearer <your_token>`

## Endpoints

### Authentication

#### Register User
```http
POST /auth/signup
```

Request body:
```json
{
    "first_name": "James",
    "last_name": "Bond",
    "phone": "+1234567892",
    "email": "guest@example.com",
    "password": "StrongPass123!",
    "role": "MANAGER"
}
```

Response:
```json
{
    "status": "success",
    "message": "User registered successfully",
    "data": {
        "id": "uuid",
        "first_name": "James",
        "last_name": "Bond",
        "email": "guest@example.com",
        "phone": "+1234567892",
        "role": "MANAGER"
    }
}
```

#### Login
```http
POST /auth/signin
```

Request body:
```json
{
    "email": "guest@example.com",
    "password": "StrongPass123!"
}
```

Response:
```json
{
    "status": "success",
    "message": "Login successful",
    "data": {
        "access_token": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
        "user": {
            "id": "uuid",
            "email": "guest@example.com",
            "role": "MANAGER"
        }
    }
}
```

### Room Management

#### Room Types

##### Create Room Type (Requires ADMIN/MANAGER Role)
```http
POST /room/create-type
```

Request body:
```json
{
    "name": "one-bedroom",
    "description": "One bedroom bnbs",
    "price_per_night": 199.99
}
```

##### List All Room Types
```http
GET /room/type/all
```

##### Get Room Type Details
```http
GET /room/type-details/{type_id}
```

#### Rooms

##### Search Available Rooms
```http
GET /room/available?check_in=2025-02-18&check_out=2025-02-22
```

Query Parameters:
- `check_in`: Check-in date (YYYY-MM-DD)
- `check_out`: Check-out date (YYYY-MM-DD)

##### Create Room (Requires ADMIN Role)
```http
POST /room/create-room
```

Request body:
```json
{
    "room_number": 123,
    "floor_number": 3,
    "room_type_id": "uuid"
}
```

##### List All Rooms
```http
GET /room/all-rooms
```

##### Get Room Details
```http
GET /room/room-details/{room_id}
```

### Reservations

#### Create Reservation
```http
POST /reservation/create-reservation
```

Request body:
```json
{
    "room_id": "uuid",
    "check_in_date": "2025-02-18",
    "check_out_date": "2025-02-22"
}
```

#### Cancel Reservation
```http
PATCH /reservation/cancel/{reservation_id}
```

Request body: none
```

#### List User Reservations
```http
GET /reservation/me
```

#### Get Reservation Details
```http
GET /reservation/reservation-details/{reservation_id}
```

### Payments

#### Create Payment
```http
POST /payments
```

Request body:
```json
{
    "reservation_id": "uuid",
    "payment_method": "CREDIT_CARD",
    "amount": 499.99
}
```

Authorization: Bearer Token required

#### Get Payment Details
```http
GET /payments/{payment_id}
```

Authorization: Bearer Token required

## Error Responses

### 400 Bad Request
```json
{
    "status": "error",
    "message": "Invalid input data",
    "errors": ["Specific validation errors"]
}
```

### 401 Unauthorized
```json
{
    "status": "error",
    "message": "Authentication required"
}
```

### 403 Forbidden
```json
{
    "status": "error",
    "message": "Insufficient permissions"
}
```

### 404 Not Found
```json
{
    "status": "error",
    "message": "Resource not found"
}
```