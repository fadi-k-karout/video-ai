# Video AI

AI-powered video processing and enhancement platform.

## Quick Start
```bash
task dev    # Start development environment
```

## Documentation
- [`docs/product.md`](docs/product.md) - Product overview and business objectives
- [`docs/tech.md`](docs/tech.md) - Technology stack and architecture
- [`docs/structure.md`](docs/structure.md) - Project structure and conventions

## Architecture
- **Backend**: Go + Gin (localhost:8080)
- **Frontend**: SvelteKit + TypeScript (localhost:5173)
- **Development**: Task + Tmux + Air hot reload

## Development Commands
```bash
task dev          # Start both services in tmux
task dev:backend  # Backend only
task dev:frontend # Frontend only
task dev:kill     # Stop all services
task build        # Production build
```

## Project Structure
```
video-ai/
├── backend/      # Go API server
├── frontend/     # SvelteKit web app
├── docs/         # Technical and product documentation
└── Taskfile.yml  # Development workflow
```

For detailed information, see the documentation files in the `docs/` directory.