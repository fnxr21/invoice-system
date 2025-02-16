# Invoice System

The Invoice System streamlines invoice creation and tracking for both paid and unpaid transactions. It manages customer data and job types, including services and equipment. Designed for efficient monitoring, the system focuses solely on generating and tracking invoices without handling payment updates.

Link

- [ERD](https://drawsql.app/teams/team-1756/diagrams/invoice)
- [Postmant](https://crimson-crater-385688.postman.co/workspace/Open~b2b49bd4-fdb3-472f-ae87-d815a933728a/collection/27541101-2afa9ac1-36c0-47dc-8996-ba565965fb0e?action=share&creator=27541101&active-environment=27541101-3f0ffd3d-2f19-42c1-a63b-666145fff426)
- [GitHub](https://github.com/fnxr21/invoice-system)

## folder structur

```plaintext
├── cmd/                      # Contains application entry points or command-line commands
│   └── cmd.go                # Main server setup and initialization
├── doc/                      # Documentation files (e.g., API specs, design notes)
│   ├── api/                  # API-related documentation
│   ├── README.md             # Documentation about Git strategy
│   └── ...
├── internal/                 # Private application code, organized by domain
│   ├── database/             # Database setup and access (e.g., connection, migrations)
│   ├── dto/                  # Data Transfer Objects (DTOs) for structuring input and output data
│   ├── handler/              # Handles HTTP requests and interacts with services
│   ├── models/               # Data models or ORM entities for database operations
│   ├── repository/           # Interfaces and implementations for data persistence
│   ├── router/               # HTTP router setup, including API route definitions
│   ├── service/              # Business logic and core application services
├── pkg/                      # Utility and reusable libraries shared across the application
│   ├── error/                # Error handling utilities
│   ├── logger/               # Logging utilities
│   ├── mysql/                # MySQL-related utilities
│   └── validate/             # Data validation utilities
├── .env                      # Environment variables configuration file
├── docker-compose.yaml       # Docker Compose file for container orchestration
├── Dockerfile                # Dockerfile for building the application container
├── main.go                   # Application entry point
├── Makefile                  # Automation tasks (e.g., build, test, deploy)
├── push.sh                   # Script for automating Git push operations
└── README.md                 # Project overview and usage guide
```

# Running in Local

Clone the repository

```cmd
git clone https://github.com/fnxr21/invoice-system
```

Go to the project repository

```cmd
cd invoice-system

```

Install dependencies

```cmd
go mod tidy
```

.env

```env
SECRET_KEY=SECRET_KEY


DB_HOST1=host.docker.internal //another setup for connection(use this if apps running in docker)
DB_HOST=127.0.0.1
APP_PORT=5001
DB_PORT=1306
DB_USER=root21!
DB_PASS=root21!Save
DB_ROOTPASSWORD=root21!SaveMain

DB_CHARNTIME=?charset=utf8mb4&parseTime=True&loc=Local


# database relation
DB_MST=invoice_mst
DB_PORT_EXPOSE=1306
DB_DEFAULTPORT=3306
```

running docker mysql db version 5.7

```cmd
docker compose up -d
```

Start the server

```cmd
go run main.go
```

## simple Unit Testing

open service

```
cd internal/service
```

running unitesting

```
got test -v
```
