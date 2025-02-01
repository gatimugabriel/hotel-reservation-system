# Hotel Reservation System

## Overview
This system is built raw Golang and PostgreSQL (Using GORM as the ORM). The system provides functionalities for room booking, user management, payments, and simple notifications.

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

## Architecture

[//]: # (![Architecture Diagram]&#40;docs/images/architecture.png&#41;)

The system follows a domain driven design (DDD) architecture with the following components and domains:
- **API Component**: Handles routing and authentication
- **User Domain**: Manages user accounts and authentication
- **Room Domain**: Handles room(and room types also) functionalities (CRUD)
- **Reservation(Booking) Domain**: Handles room reservations and availability
- **Payment Service**: payments handling 

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
│   ├── constants
│   ├── domain
│   │   ├── payment
│   │   │   ├── entity
│   │   │   └── repository
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
├── pkg
│   └── utils
│       └── input


```


## Features
- User Authentication and Authorization
- Room Availability Check
- Room Booking Management
- Payment details capturing
- Email Notifications for bookings
- Booking(Reservation) History

## Prerequisites
- Go 1.19+ (if using docker to run, no need to have this)
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
CREATE DATABASE hrs;
```

3. Configure environment variables
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Install dependencies
```bash
go mod tidy && go mod download
```

6. Start the server
```bash
make build 
make run
```

To have includes live reloading capabilities, run:
```bash
make watch 
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
OR
```bash
make docker-up
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
go test ./... -run <nameOfTest>

```