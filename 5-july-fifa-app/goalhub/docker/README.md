# Docker notes
Use `docker compose up --build` from the repository root. The backend waits for PostgreSQL health, applies Alembic migrations, seeds demo data, and starts FastAPI.
