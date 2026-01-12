# Video AI - Project Structure

## File Organization

### Root Structure
```text
video-ai/
├── backend/                     # Go backend application
├── frontend/                    # SvelteKit frontend application
├── docs/                        # Documentation (product.md, tech.md, structure.md)
├── .github/                     # GitHub templates and workflows
└── Taskfile.yml                 # Task runner configuration
```

### Backend Structure
```text
backend/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/                    # Private application code
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── handlers/
│   │   ├── health.go            # Health check handler
│   │   └── health_test.go       # Handler tests
│   ├── middleware/
│   │   └── logger.go            # Request logging, error handling
│   ├── routes/
│   │   └── routes.go            # Route definitions and setup
│   ├── services/                # Business logic layer
│   ├── models/                  # Data structures
│   └── errors/                  # Custom error types
├── pkg/                         # Public packages (reusable)
├── configs/                     # Configuration files
├── scripts/                     # Deployment and utility scripts
├── .env                         # Environment variables (local)
├── .env.example                 # Environment template
├── .air.toml                    # Hot reload configuration
├── go.mod                       # Go module definition
└── go.sum                       # Go module checksums
```

### Frontend Structure
```text
frontend/
├── src/
│   ├── lib/                     # Reusable components and utilities
│   ├── routes/                  # SvelteKit pages and layouts
│   ├── app.html                 # HTML template
│   └── app.d.ts                 # TypeScript declarations
├── static/                      # Static assets
├── package.json                 # Node.js dependencies
├── svelte.config.js             # SvelteKit configuration
├── vite.config.ts               # Vite build configuration
└── tsconfig.json                # TypeScript configuration
```

## Naming Conventions

### Go Backend
- **Packages**: lowercase, single word (`config`, `handlers`, `middleware`)
- **Files**: lowercase with underscores (`health_test.go`, `request_logger.go`)
- **Functions**: PascalCase for exported (`HealthCheck`), camelCase for private (`loadConfig`)
- **Structs**: PascalCase (`Config`, `APIError`)
- **Constants**: PascalCase or UPPER_CASE (`RequestIDKey`, `DEFAULT_PORT`)
- **Interfaces**: PascalCase with -er suffix (`Logger`, `Handler`)

### Frontend
- **Components**: PascalCase (`.svelte` files)
- **Routes**: lowercase with hyphens (`/api/health`, `/user-profile`)
- **Utilities**: camelCase (`formatDate`, `apiClient`)
- **Types**: PascalCase (`User`, `VideoProcessingRequest`)

## Import Patterns

### Go Imports
```go
// Standard library first
import (
    "context"
    "log/slog"
    "net/http"
)

// Third-party packages
import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// Local packages (absolute from module root)
import (
    "videoai/internal/config"
    "videoai/internal/handlers"
)
```

### TypeScript Imports
```typescript
// Relative imports for local files
import { apiClient } from '$lib/api';
import type { User } from '$lib/types';

// External packages
import { writable } from 'svelte/store';
```

## Architectural Decisions

### Backend Patterns
- **Layered Architecture**: handlers → services → models
- **Dependency Injection**: Pass dependencies through constructors
- **Error Handling**: Custom error types with HTTP status codes
- **Middleware Chain**: logger → CORS → auth → error handler
- **Configuration**: Environment-based with struct validation
- **Testing**: Table-driven tests, behavior over implementation

### API Design
- **RESTful endpoints**: `/api/v1/resource`
- **JSON responses**: Consistent structure with error codes
- **Request IDs**: UUID for request tracing
- **Status codes**: Standard HTTP codes (200, 400, 404, 500)
- **CORS**: Configured for frontend origins

### Frontend Patterns
- **Component Structure**: Single-file components (.svelte)
- **State Management**: Svelte stores for global state
- **API Integration**: Centralized API client
- **Type Safety**: TypeScript throughout
- **Routing**: File-based routing with SvelteKit

## Code Organization Rules

### Backend
- Keep handlers thin - delegate to services
- Business logic in services, not handlers
- Models are data structures only
- Middleware for cross-cutting concerns
- Tests alongside implementation files

### Frontend
- Components in `src/lib/components/`
- Utilities in `src/lib/utils/`
- Types in `src/lib/types/`
- API client in `src/lib/api/`
- Pages in `src/routes/`

## Development Workflow
- Feature branches from main
- One feature per branch
- Tests required for new code
- Documentation updates with features
- GitHub issues for task tracking