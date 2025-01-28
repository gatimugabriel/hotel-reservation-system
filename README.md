# Hotel Reservation System

## Overview
This is a comprehensive hotel reservation system built with Go and PostgreSQL. The system provides functionalities for room booking, user management, payments processing, and notifications.

## Table of Contents
- [Hotel Reservation System](#hotel-reservation-system)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [Architecture](#architecture)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Setup \& Installation](#setup--installation)
    - [Manual Setup](#manual-setup)
    - [Docker Setup](#docker-setup)
  - [API Documentation](#api-documentation)
  - [Database Schema](#database-schema)
  - [Development Guidelines](#development-guidelines)
  - [Testing](#testing)
  - [Deployment](#deployment)

## Architecture
![Architecture Diagram](docs/images/architecture.png)

The system follows a domain driven design (DDD) architecture with the following components:
- **API**: Handles routing and authentication
- **User Service**: Manages user accounts and authentication
- **Booking Service**: Handles room reservations and availability
- **Hotel Service**: Handles hotel
- **Payment Service**: Processes payments 
- **Notification Service**: Manages email and SMS(to be added later) notifications
- **Room Inventory Service**: Manages room types and availability

## Code Structure
As explained above, the system adopts a domain driven design where each domain has its own entities(model(s) definition), repository(for database interaction & data persistence), service(business logic) 
. I adopted this structure as:
  1. It scales easily
  2. Easier to debug
  3. Not complex enough to be like them, but somehow separates concern like microservices

The structure
```plaintext
.
├── cmd
│   └── app
├── diagrams
├── docs
├── internal
│   ├── api
│   │   ├── handlers
│   │   ├── middleware
│   │   └── router
│   ├── config
│   ├── domain
│   │   ├── hotel
│   │   │   ├── entity
│   │   │   ├── repository
│   │   │   └── services
│   │   ├── payment
│   │   ├── reservation
│   │   │   ├── entity
│   │   │   ├── repository
│   │   │   └── services
│   │   ├── room
│   │   │   ├── entity
│   │   │   ├── repository
│   │   │   └── services
│   │   └── user
│   │       ├── entity
│   │       ├── repository
│   │       └── services
│   ├── infrastructure
│   │   └── database
│   │       ├── raw_queries
│   │       └── sql
│   └── server
│       └── httpServer
└── pkg
    └── utils

```


## Features
- User Authentication and Authorization
- Room Availability Check
- Room Booking Management
- Payment Processing
- Email/SMS Notifications
- Admin Dashboard
- Booking History
- Room Inventory Management

## Prerequisites
- Go 1.19+
- PostgreSQL 14+
- Docker & Docker Compose (for containerized setup)
- Make (optional, for using Makefile commands)

## Setup & Installation

### Manual Setup
1. Clone the repository
```bash
git clone https://github.com/gatimugabriel/hotel-reservation-system
cd hotel-reservation-system
```

2. Set up the database
```bash
psql -U postgres
CREATE DATABASE hotel_reservation;
```

3. Configure environment variables
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Install dependencies
```bash
go mod download
```

5. Run migrations
```bash
make migrate-up
```

6. Start the server
```bash
make run
```

### Docker Setup
1. Clone the repository
```bash
git clone https://github.com/gatimugabriel/hotel-reservation-system
cd hotel-reservation-system
```

2. Build and run with Docker Compose
```bash
docker-compose up --build
```

The application will be available at `http://localhost:8080`

## API Documentation
Detailed API documentation can be found in [docs/api.md](docs/api.md)

## Database Schema
Database schema and relationships are documented in [docs/database.md](docs/database.md)

## Development Guidelines
Please refer to [docs/development.md](docs/development.md) for coding standards, best practices, and contribution guidelines.

## Testing
```bash
# Run all tests
make test

# Run specific tests
go test ./... -run TestNameHere

# Run with coverage
make test-coverage
```

## Deployment
Deployment instructions and considerations can be found in [docs/deployment.md](docs/deployment.md)