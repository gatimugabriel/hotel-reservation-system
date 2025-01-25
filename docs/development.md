# Development Guidelines

## Code Structure

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

## Coding Standards

### Go Specific Guidelines

1. Follow the official Go style guide
2. Use meaningful variable and function names
3. Keep functions small and focused
4. Use proper error handling
5. Document public APIs and complex logic
6. Use interfaces for better testability
7. Follow the SOLID principles

### Example

```go
// Good Example
type BookingService interface {
    CreateBooking(ctx context.Context, booking *models.Booking) error
    GetBookingByID(ctx context.Context, id string) (*models.Booking, error)
}

type bookingService struct {
    repo    repository.BookingRepository
    logger  *logger.Logger
    metrics *metrics.Metrics
}

func NewBookingService(repo repository.BookingRepository, logger *logger.Logger, metrics *metrics.Metrics) BookingService {
    return &bookingService{
        repo:    repo,
        logger:  logger,
        metrics: metrics,
    }
}

func (s *bookingService) CreateBooking(ctx context.Context, booking *models.Booking) error {
    if err := booking.Validate(); err != nil {
        return fmt.Errorf("invalid booking data: %w", err)
    }

    if err := s.repo.Create(ctx, booking); err != nil {
        s.logger.Error("failed to create booking", "error", err)
        return fmt.Errorf("failed to create booking: %w", err)
    }

    s.metrics.BookingCreated.Inc()
    return nil
}
```

## Error Handling

1. Use custom error types for domain-specific errors
2. Always wrap errors with context
3. Log errors at the appropriate level
4. Return appropriate HTTP status codes

```go
type ErrorResponse struct {
    Status  string `json:"status"`
    Code    string `json:"code"`
    Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, err error) {
    var resp ErrorResponse
    switch {
    case errors.Is(err, ErrNotFound):
        resp = ErrorResponse{
            Status:  "error",
            Code:    "NOT_FOUND",
            Message: err.Error(),
        }
        w.WriteHeader(http.StatusNotFound)
    case errors.Is(err, ErrInvalidInput):
        resp = ErrorResponse{
            Status:  "error",
            Code:    "INVALID_INPUT",
            Message: err.Error(),
        }
        w.WriteHeader(http.StatusBadRequest)
    default:
        resp = ErrorResponse{
            Status:  "error",
            Code:    "INTERNAL_ERROR",
            Message: "An internal error occurred",
        }
        w.WriteHeader(http.StatusInternalServerError)
    }

    json.NewEncoder(w).Encode(resp)
}
```

## Testing

1. Write unit tests for all business logic
2. Use table-driven tests where appropriate
3. Use mocks for external dependencies
4. Aim for high test coverage
5. Include integration tests for critical paths

```go
func TestCreateBooking(t *testing.T) {
    tests := []struct {
        name    string
        booking *models.Booking
        wantErr bool
    }{
        {
            name: "valid booking",
            booking: &models.Booking{
                UserID:       "user-1",
                RoomID:       "room-1",
                CheckInDate:  time.Now(),
                CheckOutDate: time.Now().AddDate(0, 0, 1),
            },
            wantErr: false,
        },
        // Add more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockRepo := mock.NewMockBookingRepository()
            service := NewBookingService(mockRepo, logger.New(), metrics.New())

            err := service.CreateBooking(context.Background(), tt.booking)
            if (err != nil) != tt.wantErr {
                t.Errorf("CreateBooking() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

## Logging

1. Use structured logging
2. Include relevant context
3. Use appropriate log levels
4. Include request ID in logs

```go
func (h *bookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    logger := h.logger.With(
        "request_id", middleware.GetRequestID(ctx),
        "handler", "CreateBooking",
    )

    var req CreateBookingRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        logger.Error("failed to decode request", "error", err)
        WriteError(w, ErrInvalidInput)
        return
    }

    logger.Info("creating booking", "user_id", req.UserID)
    // ... rest of the handler
}
```

## API Design

1. Follow RESTful conventions
2. Use proper HTTP methods and status codes
3. Version your APIs
4. Include proper validation
5. Document APIs using OpenAPI/Swagger

## Security

1. Use HTTPS
2. Implement proper authentication and authorization
3. Sanitize user input
4. Use secure password hashing
5. Implement rate limiting
6. Use secure headers

## Performance

1. Use connection pooling for database
2. Implement caching where appropriate
3. Use proper indexing
4. Monitor and optimize database queries
5. Implement proper timeouts

## Deployment

1. Use Docker for containerization
2. Implement health checks
3. Use proper environment variables
4. Implement graceful shutdown
5. Use proper logging and monitoring

## Contribution Guidelines

1. Fork the repository
2. Create a feature branch
3. Write tests for new features
4. Follow coding standards
5. Submit a pull request with proper description

## Code Review Guidelines

1. Check for proper error handling
2. Verify test coverage
3. Review security implications
4. Check for performance issues
5. Ensure documentation is updated