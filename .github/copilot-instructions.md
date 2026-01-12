# GitHub Copilot - Project Context

This project uses:
- Go 1.25.5 with Gin framework for backend API
- SvelteKit with TypeScript for frontend
- Task runner for development workflow (`task dev`)
- Structured logging with slog and request ID tracking
- Layered architecture: handlers → services → models

Key files:
- docs/product.md - Business requirements
- docs/tech.md - Technical stack and constraints  
- docs/structure.md - Code organization patterns
- Taskfile.yml - Development commands

Development workflow:
```bash
task dev        # Start both backend and frontend
task dev:kill   # Stop development environment
task build      # Production build
```

Follow the patterns in docs/structure.md for naming conventions and project organization.