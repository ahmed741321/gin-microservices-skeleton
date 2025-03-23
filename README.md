# gin-microservices-app

## Overview

This project is a microservices architecture built using the Gin framework in Go. It consists of an API service and a worker service, allowing for efficient handling of HTTP requests and background tasks.

## Project Structure

```
gin-microservices-app
├── cmd
│   ├── api          # Entry point for the API service
│   └── worker       # Entry point for the worker service
├── configs          # Configuration files
├── internal         # Internal application logic
│   ├── api          # API related components
│   ├── worker       # Worker related components
│   └── models       # Data models
├── pkg             # Shared packages
│   ├── logger       # Logging utilities
│   ├── database     # Database connection and queries
│   └── utils        # Utility functions
├── scripts          # Scripts for automation
├── test             # Unit tests
├── go.mod           # Go module definition
└── go.sum           # Go module checksums
```

## Getting Started

### Prerequisites

- Go 1.16 or later
- A working Go environment

### Installation

1. Clone the repository:

   ```
   git clone <repository-url>
   cd gin-microservices-app
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

### Running the API Service

To run the API service, navigate to the `cmd/api` directory and execute:

```
go run main.go
```

### Running the Worker Service

To run the worker service, navigate to the `cmd/worker` directory and execute:

```
go run main.go
```

### Database Migration

Run the migration script to set up the database:

```
sh scripts/migrate.sh
```

## Internal API Directory

The `internal/api` directory is responsible for the core logic of the API service. It contains the following components:

1. **Routes**:

   - Defines the API endpoints and maps them to their respective handlers or controllers.
   - Example:
     ```go
     func SetupRoutes(router *gin.Engine) {
         router.GET("/health", healthCheckHandler)
         router.POST("/users", createUserHandler)
         // Other routes...
     }
     ```

2. **Controllers**:

   - Contains the logic for handling requests and sending responses. Controllers interact with services or models to process data.
   - Example:
     ```go
     func createUserHandler(c *gin.Context) {
         var user User
         if err := c.ShouldBindJSON(&user); err != nil {
             c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
             return
         }
         // Call service to save user...
         c.JSON(http.StatusCreated, user)
     }
     ```

3. **Middleware**:

   - Handles tasks like authentication, logging, or request validation.
   - Example:
     ```go
     func AuthMiddleware() gin.HandlerFunc {
         return func(c *gin.Context) {
             token := c.GetHeader("Authorization")
             if token == "" {
                 c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
                 return
             }
             // Validate token...
             c.Next()
         }
     }
     ```

4. **Services**:

   - Encapsulates business logic and interacts with models or external systems (e.g., databases, APIs).
   - Example:
     ```go
     func CreateUser(user User) error {
         return database.Save(&user).Error
     }
     ```

5. **Utilities**:
   - Contains helper functions specific to the API service, such as response formatting or error handling.

## Usage

- The API service handles HTTP requests and routes them to the appropriate controllers.
- The worker service processes background tasks as defined in the tasks directory.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License.
# gin-microservices-skeleton
