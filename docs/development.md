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
└── tmp

38 directories


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


## Contribution Guidelines

1. Fork the repository
2. Create a feature branch
3. Write tests for new features
4. Follow coding standards
5. Submit a pull request with proper description