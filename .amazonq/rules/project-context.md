# Amazon Q Developer - Project Context

## Project Overview
This is the Video AI project - a web application for AI-powered video processing and enhancement.

## Key Documentation Files
- `docs/product.md` - Product requirements and business objectives
- `docs/tech.md` - Technology stack and architecture decisions  
- `docs/structure.md` - Project structure and coding conventions
- `Taskfile.yml` - Development workflow and build commands

## Development Commands
- `task dev` - Start development environment (backend + frontend in tmux)
- `task build` - Build for production
- `task dev:kill` - Stop development environment

## Architecture
- **Backend**: Go + Gin (port 8080)
- **Frontend**: SvelteKit + TypeScript (port 5173)
- **Development**: Task runner + Tmux + Air hot reload + Bun

## Coding Standards
- Follow patterns in `docs/structure.md`
- Use structured logging with slog
- Implement layered architecture (handlers → services → models)
- Write behavior-focused tests