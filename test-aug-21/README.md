# Arc — developer dashboard

A polished, dark-first personal developer workspace. Arc tracks repositories, coding momentum, priorities, notes, and Pomodoro sessions in one focused dashboard.

## Run locally

```bash
npm start
```

Then open [http://localhost:5174](http://localhost:5174). The application uses an Express API and Node's built-in SQLite module; no separate database setup is required. The database is automatically created at `data/arc.db` and initialized with believable starter tasks and a note on first launch.

## Development

```bash
npm run dev
npm run check
```

## Architecture

- `server/index.js` — lightweight REST API and SQLite persistence for tasks, notes, and focus sessions.
- `public/index.html` — dashboard structure and accessible interactive controls.
- `public/app.js` — UI behavior, keyboard commands, API client, timer, and drag ordering.
- `public/style.css` — responsive dark visual system.

Repository and contribution data are seeded client-side, behind a presentation layer designed to be replaced by GitHub-backed services later.
