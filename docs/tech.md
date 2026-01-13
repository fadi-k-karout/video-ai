# Video AI - Technology Stack

## Core Technologies

### Backend Stack
- **Language**: Go 1.25.5 (chosen for performance, concurrency, and strong typing)
- **Web Framework**: Gin Gonic (high-performance HTTP router with middleware support)
- **Logging**: slog (Go's structured logging for observability)
- **Configuration**: Environment-based config with struct validation
- **Development**: Air for hot reloading, built-in testing framework

### Frontend Stack
- **Framework**: SvelteKit (modern, performant, with SSR capabilities)
- **Language**: TypeScript (type safety and developer experience)
- **Build Tool**: Vite (fast development and optimized builds)
- **Styling**: CSS with component-scoped styles
- **Package Manager**: Bun (fast package management and runtime)

### Development Tools
- **Task Runner**: Taskfile.yml for cross-platform build automation and development workflow
- **Process Management**: Tmux for split-pane development environment
- **Hot Reload**: Air for Go backend development
- **Package Manager**: Bun for frontend dependencies and runtime
- **Version Control**: Git with GitHub for collaboration
- **IDE Integration**: VS Code with Go and Svelte extensions

## Architecture Patterns

### Backend Architecture
- **Layered Architecture**: Clear separation of concerns (handlers → services → models)
- **Dependency Injection**: Constructor-based dependency management
- **Middleware Pattern**: Cross-cutting concerns (logging, CORS, error handling)
- **Repository Pattern**: Data access abstraction (planned)
- **Clean Architecture**: Domain-driven design principles

### API Design
- **RESTful APIs**: Standard HTTP methods and status codes
- **JSend Compliance**: Structured response format (status, message, data, code)
- **Request Tracing**: UUID-based request correlation
- **Error Handling**: Layered error system with internal/client message separation
- **Versioning**: API versioning strategy (/api/v1/)
- **Security**: No sensitive data in client-facing error messages

### Frontend Architecture
- **Component-Based**: Reusable Svelte components
- **State Management**: Svelte5 runes for state
- **Type Safety**: TypeScript throughout the application
- **API Integration**: Centralized HTTP client with error handling
- **Responsive Design**: Mobile-first approach

## Technical Constraints

### Performance Requirements
- **Response Time**: < 200ms for API endpoints
- **Throughput**: Support 1000+ concurrent requests
- **Memory Usage**: Efficient memory management for video processing
- **Scalability**: Horizontal scaling capability

### Security Requirements
- **CORS**: Properly configured for frontend origins
- **Input Validation**: Server-side validation for all inputs
- **Error Handling**: No sensitive information in error responses
- **Logging**: No PII or sensitive data in logs
- **Environment Isolation**: Separate configs for dev/staging/prod

### Development Constraints
- **Go Version**: Minimum Go 1.21 for slog support
- **Browser Support**: Modern browsers (ES2020+)
- **Mobile Compatibility**: Responsive design for mobile devices
- **Deployment**: Container-ready for cloud deployment

## Libraries and Dependencies

### Backend Dependencies
```go
// Core framework
github.com/gin-gonic/gin          // HTTP web framework
github.com/gin-contrib/cors       // CORS middleware

// Utilities
github.com/google/uuid            // UUID generation

// Future additions
// Database: github.com/lib/pq (PostgreSQL)
// Validation: github.com/go-playground/validator
// Testing: github.com/stretchr/testify
```

### Frontend Dependencies
```json
{
  "@sveltejs/kit": "^2.49.1",
  "typescript": "^5.9.3",
  "vite": "^7.2.6"
}
```

## Development Workflow

### Local Development
**Primary Development Command:**
```bash
task dev  # Starts both backend and frontend in split tmux session
```

**Individual Services:**
```bash
task dev:backend   # Backend only (Air hot reload)
task dev:frontend  # Frontend only (Bun dev server)
```

**Development Environment:**
- **Backend**: Air hot reload on file changes (port 8080)
- **Frontend**: Bun dev server with HMR (port 5173)
- **Tmux Integration**: Split-pane view with separate logs
- **API Access**: `http://localhost:8080/api/`
- **Environment**: `.env` file for local configuration

**Session Management:**
```bash
task dev        # Start development environment
task dev:kill   # Stop tmux session and all services
```

### Testing Strategy
- **Unit Tests**: Go's built-in testing framework
- **HTTP Tests**: `httptest` package for API testing
- **Frontend Tests**: Vitest for component and unit testing
- **Integration Tests**: End-to-end API testing
- **Behavior Testing**: Focus on external behavior over implementation

### Build and Deployment
**Production Build:**
```bash
task build  # Builds both backend and frontend
```

**Individual Builds:**
```bash
task build:backend   # Go binary to dist/server
task build:frontend  # SvelteKit static build
```

**Build Outputs:**
- **Backend**: Single Go binary (`dist/server`)
- **Frontend**: Static files for CDN deployment
- **Containerization**: Docker for consistent deployment
- **CI/CD**: GitHub Actions for automated testing and deployment

## Monitoring and Observability

### Logging Strategy
- **Structured Logging**: JSON format with slog
- **Request Tracing**: UUID-based correlation across services
- **Log Levels**: Debug, Info, Warn, Error with environment-based filtering
- **Centralized Logging**: Grafana Loki for log aggregation

### Metrics and Monitoring
- **Application Metrics**: Request rate, response time, error rate
- **System Metrics**: CPU, memory, disk usage
- **Business Metrics**: User engagement, processing success rate
- **Alerting**: Grafana alerts for critical issues

## Future Technical Roadmap

### Database Integration
- **PostgreSQL**: Primary database for user data and metadata
- **Redis**: Caching and session management
- **Object Storage**: AWS S3 or compatible for video files

### Advanced Features
- **Authentication**: JWT-based user authentication
- **Rate Limiting**: API rate limiting and quota management
- **WebSockets**: Real-time processing updates
- **Message Queue**: Background job processing
- **Microservices**: Service decomposition for scaling