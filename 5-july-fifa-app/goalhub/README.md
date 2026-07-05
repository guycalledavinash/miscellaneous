# GoalHub

GoalHub is a production-oriented, containerized 3-tier international soccer dashboard.

## Stack
- Frontend: React, TypeScript, Vite, Material UI, Recharts
- Backend: FastAPI, SQLAlchemy, Pydantic, Alembic, JWT, structured logging
- Database: PostgreSQL with a named Docker volume

## Quick start
```bash
docker compose up --build
```

Open the frontend at http://localhost:3000 and the API docs at http://localhost:8000/docs.

Demo users:
- `admin` / `admin123` can create, edit, and delete.
- `viewer` / `viewer123` can read.

## Development
```bash
cd backend && python -m venv .venv && source .venv/bin/activate && pip install -r requirements-dev.txt
ruff check app && black --check app && pytest
```

```bash
cd frontend && npm install && npm run lint && npm run build
```

## Seed data
`scripts/seed.py` creates 20 teams, 200 players, 50 matches, four competitions, and demo users. Docker Compose runs it after migrations.

See `docs/architecture.md` for Mermaid architecture and ER diagrams, and `docs/api.md` for endpoint documentation.
