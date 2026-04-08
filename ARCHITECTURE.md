# Calm-Map Architecture Improvement Plan

## Context

Calm-Map is currently a pure frontend app: vanilla HTML/CSS/JS, a static `points.json` file, and no server. Every backlog feature — comments, favorites, upvotes, user submissions, journal entries — requires server-side state. The team is primarily backend Go developers, so introducing a Go HTTP server is a natural fit. This plan adds a Go backend while keeping the frontend vanilla JS (no new frontend tooling) and positions the project to deliver its full backlog.

---

## Recommended Architecture

**Go HTTP server (Chi) + PostgreSQL + PostGIS**, with frontend and backend deployed separately.

```
Browser ──► Caddy (localhost:3000 / GitHub Pages)
               ├── /              → serves web/ static files
               └── /api/*         → proxies to Go backend

            Go server (Chi) (localhost:8080 / Railway)
               └── /api/v1/...    → JSON REST API
                        │
                        └── PostgreSQL + PostGIS (Railway)
```

---

## Key Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| Web framework | **Chi** | Wraps `net/http`; handlers are plain `http.HandlerFunc`; no new context types to learn |
| Database | **PostgreSQL + PostGIS** | Backlog includes polygons/lines; PostGIS handles spatial queries natively; SQLite write concurrency is a problem for social features |
| Query layer | **sqlc** (or pgx/v5 directly) | No GORM; write SQL, get type-safe Go; avoids magic and N+1 surprises |
| Static serving | **Caddy** (dev) / **GitHub Pages** (prod) | No embedding; frontend deploys independently; Caddy proxies API in dev to avoid CORS |
| Frontend tooling | **None** | Team is not frontend-heavy; vanilla JS with ES modules is sufficient |
| Auth (v1) | **Session tokens + HTTP-only cookie** | Simple, stateful, no JWT key rotation complexity |

---

## Directory Structure

Code is grouped by **domain**, not by layer. No `handlers/`, `models/`, or `services/` packages — those are kitchen-drawer anti-patterns. Each domain package owns its handlers, types, and store queries together. Dependencies are passed via constructor injection; no globals.

```
calm-map/
├── cmd/
│   ├── server/main.go         # Entry point — thin: config, deps, start, shutdown
│   └── seed/main.go           # One-time: imports points.json into DB
├── internal/
│   ├── places/
│   │   ├── handler.go         # Chi handlers for /api/v1/places
│   │   ├── service.go         # Business logic (filtering, approval)
│   │   └── store.go           # SQL queries (sqlc-generated or pgx direct)
│   ├── comments/
│   │   ├── handler.go
│   │   └── store.go
│   ├── favorites/
│   │   ├── handler.go
│   │   └── store.go
│   ├── votes/
│   │   ├── handler.go
│   │   └── store.go
│   └── platform/
│       ├── database/db.go     # pgx pool init — injected into stores
│       └── config/config.go   # Reads env vars (DATABASE_URL, PORT, etc.)
├── migrations/
│   ├── 001_create_places.sql
│   ├── 002_create_users.sql
│   ├── 003_create_comments.sql
│   ├── 004_create_favorites.sql
│   └── 005_create_votes.sql
├── web/                       # Static frontend (deployed separately to GitHub Pages)
│   ├── index.html
│   ├── app.js                 # JS extracted from index.html
│   ├── style.css
│   └── docs/
├── design/                    # Architecture Decision Records (ADRs)
│   └── ADR-001-go-backend.md
├── docker-compose.yml
├── Dockerfile.dev             # Go + Air, for local development
├── Dockerfile                 # Multi-stage production build
├── Caddyfile.dev              # Caddy config: serves web/, proxies /api/* to Go
├── Makefile                   # make dev / make migrate / make clean
├── .air.toml                  # Air live-reload config
├── sqlc.yaml
├── go.mod
└── .env.example
```

### main.go pattern

`cmd/server/main.go` stays thin. Dependencies are initialised once and injected via constructors — no globals.

```go
func main() {
    if err := run(context.Background(), os.Args, os.Environ()); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func run(ctx context.Context, args []string, env []string) error {
    // 1. Parse config from env
    // 2. Init pgx pool — pass to store constructors
    // 3. Build Chi router — pass stores/services to handlers
    // 4. Start HTTP server
    // 5. Block on shutdown signal
    return nil
}
```

---

## Database Schema (core tables)

```sql
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE places (
    id            BIGSERIAL PRIMARY KEY,
    name          TEXT NOT NULL,
    address       TEXT,
    vibe_type     TEXT,
    cost          TEXT,
    accessibility TEXT,
    image_url     TEXT,
    location      GEOMETRY(Point, 4326) NOT NULL,
    geometry      GEOMETRY,              -- future: polygons/lines
    hours         JSONB,
    submitted_by  BIGINT REFERENCES users(id),
    approved      BOOLEAN NOT NULL DEFAULT false,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX places_location_idx ON places USING GIST (location);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    handle TEXT UNIQUE NOT NULL,
    email  TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE comments (
    id BIGSERIAL PRIMARY KEY,
    place_id BIGINT NOT NULL REFERENCES places(id) ON DELETE CASCADE,
    user_id  BIGINT NOT NULL REFERENCES users(id),
    body     TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE favorites (
    user_id  BIGINT NOT NULL REFERENCES users(id),
    place_id BIGINT NOT NULL REFERENCES places(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, place_id)
);

CREATE TABLE votes (
    user_id  BIGINT NOT NULL REFERENCES users(id),
    place_id BIGINT NOT NULL REFERENCES places(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, place_id)
);
```

---

## API Design (`/api/v1/`)

```
GET  /places              ?vibe_type= &cost= &accessibility= &q= &bbox=
GET  /places/:id          (includes vote_count, comments)
POST /places              (auth required; sets approved=false for non-admin)

GET  /places/:id/comments
POST /places/:id/comments  (auth required)
DELETE /comments/:id       (auth required, own comment)

GET    /users/me/favorites  (auth required)
PUT    /places/:id/favorite (auth required)
DELETE /places/:id/favorite (auth required)

PUT    /places/:id/vote     (auth required)
DELETE /places/:id/vote     (auth required)
```

---

## Frontend Changes (minimal)

Only two changes to `index.html` at migration time:

1. `fetch("points.json")` → `fetch("/api/v1/places")`
2. Field name: `vibe-type` → `vibe_type` (one line in marker creation loop)

Optionally: extract the inline `<script>` block to `web/app.js` (quality-of-life, not required).

No build tooling added. No npm.

---

## Data Migration

`cmd/seed/main.go` — one-time script that:
- Reads `points.json`
- Inserts each record into `places` with `approved = true`
- Converts `lat`/`lng` to PostGIS point: `ST_SetSRID(ST_MakePoint(lng, lat), 4326)`
- Is idempotent (skip existing rows by name)

---

## Implementation Phases

Each phase ships independently:

| Phase | Work | Outcome |
|-------|------|---------|
| 1 | Go server with CORS + health check; move JS to `web/app.js` | Backend running, frontend still loads from file or GitHub Pages |
| 2 | Places API + DB + seed script | Places load from PostgreSQL |
| 3 | Comments API + popup UI | Users can comment |
| 4 | Auth + favorites + votes | Social features live |
| 5 | User submissions + admin moderation | Community contributions |

---

## Local Development

One command to start everything:

```bash
make dev      # starts db + api + frontend
make migrate  # run migrations (first time, or after adding new ones)
make clean    # tear down everything including volumes
```

### docker-compose.yml

```yaml
services:
  db:
    image: postgis/postgis:17-3.4
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: calm_map
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 10s
      timeout: 5s
      retries: 5

  api:
    build:
      context: .
      dockerfile: Dockerfile.dev
    environment:
      DATABASE_URL: postgres://postgres:postgres@db:5432/calm_map?sslmode=disable
    ports:
      - "8080:8080"
    volumes:
      - .:/app
      - /app/tmp
    depends_on:
      db:
        condition: service_healthy

  frontend:
    image: caddy:latest
    ports:
      - "3000:3000"
    volumes:
      - ./web:/srv
      - ./Caddyfile.dev:/etc/caddy/Caddyfile
    depends_on:
      - api

  migrate:
    image: migrate/migrate
    volumes:
      - ./migrations:/migrations
    command: -path=/migrations/ -database postgres://postgres:postgres@db:5432/calm_map?sslmode=disable up
    depends_on:
      db:
        condition: service_healthy
    profiles:
      - migrate

volumes:
  postgres_data:
```

### Dockerfile.dev (Go with Air live reload)

```dockerfile
FROM golang:1.23-alpine
WORKDIR /app
RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]
```

### Caddyfile.dev (static frontend + API proxy)

```
:3000 {
  root * /srv
  file_server
  reverse_proxy /api/* http://api:8080
  log { output stdout }
}
```

Caddy proxies `/api/*` to the Go backend — no CORS issues in local dev. The same CORS middleware in Chi handles production (where frontend and backend are on different domains).

### Key DX properties

- **Frontend changes** (HTML/CSS/JS): instant — Caddy serves directly from the `web/` volume mount
- **Backend changes** (Go): Air rebuilds in ~1–2s, no container restart needed
- **Schema changes**: add a migration file, run `make migrate`
- **First-time setup**: `cp .env.example .env && make dev && make migrate`

---

## Deployment

Frontend and backend are deployed separately. No embedding.

### Recommended Stack

| Layer | Service | Cost |
|-------|---------|------|
| Frontend (static files) | **GitHub Pages** | Free |
| Backend (Go) | **Railway** | $5/month |
| Database | **Railway PostgreSQL** (PostGIS template) | Included in $5/month |

Railway keeps backend + database in one dashboard with a PostGIS-ready template. GitHub Pages is zero-configuration for a team already on GitHub.

**Alternative:** Cloudflare Pages (frontend, unlimited bandwidth) + Fly.io (backend + managed Postgres with PostGIS, ~$10–30/month, pay-as-you-go, global distribution).

### Frontend (GitHub Pages)

No build step. Push `web/index.html`, `web/style.css`, and `web/docs/` to the repo. GitHub Pages serves them directly. Point `API_BASE_URL` in `app.js` to the Railway backend URL.

```
https://your-org.github.io/calm-map/  →  serves index.html, style.css, docs/
```

### Backend (Railway)

```dockerfile
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o calm-map ./cmd/server

FROM alpine:3.20
COPY --from=builder /app/calm-map /usr/local/bin/calm-map
EXPOSE 8080
CMD ["calm-map"]
```

Railway detects the Dockerfile automatically on push. Set `DATABASE_URL` as an environment variable in the Railway dashboard.

### CORS

Because frontend and backend are on different domains, add CORS middleware to Chi:

```go
import "github.com/go-chi/cors"

r.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"https://your-org.github.io"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
    AllowCredentials: true, // required for session cookies
}))
```

This is ~10 lines and is the only infrastructure change required for the separated deployment.

### Migrations

[golang-migrate](https://github.com/golang-migrate/migrate) CLI with plain `.sql` files in `migrations/`. Run manually as a deployment step:

```
migrate -database $DATABASE_URL -path migrations up
```

---

## Critical Files to Modify

- `index.html` — change `fetch("points.json")` call and `vibe-type` field name
- `points.json` — source of truth for seed data; every field maps to a `places` column
- `.agents/architecture.md` — must be updated to reflect Go + PostgreSQL architecture
- `.agents/workflows.md` — needs "Running the Go server" and "Running migrations" sections
- `TASKS.md` — each backlog item maps to a specific DB table + handler

---

## Verification

1. `go build ./...` — compiles cleanly
2. `go run ./cmd/seed/main.go` — inserts 10 places from `points.json`
3. `curl /api/v1/places` — returns 10 places as JSON
4. Open browser → map loads with markers from API (not local file)
5. Apply a filter → query params sent to server, filtered results returned
6. `POST /api/v1/places/:id/comments` — comment persists across page reload
