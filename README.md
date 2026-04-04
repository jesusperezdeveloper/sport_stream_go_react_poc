# SportStream — Go + React POC

Sports content management platform with a Go REST API backend and React dashboard frontend.

## Architecture

```
┌─────────────────────┐     REST/JSON     ┌──────────────────────┐
│   React Dashboard   │ ◄──────────────► │    Go API Server     │
│   (Vite + TS + TW4) │                   │   (net/http stdlib)  │
│   Port 3000/5173    │                   │     Port 8080        │
└─────────────────────┘                   └──────────────────────┘
```

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go 1.23+, net/http (enhanced routing), slog, uuid |
| Frontend | React 19, Vite 6, TypeScript, Tailwind CSS v4 |
| Data fetching | TanStack Query v5 |
| Charts | Recharts |
| Icons | Lucide React |
| Storage | In-memory (no database required) |
| Deploy | Docker multi-stage builds |

## Quick Start

### Backend

```bash
cd sportstream-api
go mod tidy
go run ./cmd/server
# API running at http://localhost:8080
```

Verify:
```bash
curl http://localhost:8080/api/v1/health
curl http://localhost:8080/api/v1/dashboard/summary
```

### Frontend

```bash
cd sportstream-dashboard
npm install
VITE_API_URL=http://localhost:8080 npm run dev
# Dashboard at http://localhost:5173
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/v1/health` | Health check + version |
| GET | `/api/v1/clubs` | List all clubs |
| GET | `/api/v1/clubs/{id}` | Club detail |
| POST | `/api/v1/clubs` | Create club |
| PUT | `/api/v1/clubs/{id}` | Update club |
| DELETE | `/api/v1/clubs/{id}` | Delete club |
| GET | `/api/v1/clubs/{id}/streams` | Club streams |
| GET | `/api/v1/streams` | List streams (filter: `?status=live&type=vod`) |
| GET | `/api/v1/streams/{id}` | Stream detail |
| POST | `/api/v1/streams` | Create stream |
| PUT | `/api/v1/streams/{id}` | Update stream |
| PATCH | `/api/v1/streams/{id}/status` | Change stream status |
| GET | `/api/v1/events` | List events (filter: `?status=upcoming&sport=football`) |
| GET | `/api/v1/events/{id}` | Event detail |
| POST | `/api/v1/events` | Create event |
| GET | `/api/v1/events/upcoming` | Upcoming events |
| GET | `/api/v1/dashboard/summary` | Dashboard aggregation |

## Seed Data

The API starts with preloaded data:

- **5 clubs**: S.S. Lazio, Sevilla FC, Lega Volley, SuperTennix, FIBA Europe
- **10 streams**: Mix of live, scheduled, VOD, highlights
- **6 events**: Mix of upcoming, live, completed

## Project Structure

```
sportstream-api/
├── cmd/server/              # Entrypoint
├── internal/
│   ├── domain/              # Entities + repository interfaces
│   ├── application/         # Business logic services
│   └── infrastructure/
│       ├── http/            # Router, handlers, middleware
│       ├── persistence/     # In-memory repositories + seed
│       └── config/          # Environment config
└── deployments/             # Dockerfile + docker-compose

sportstream-dashboard/
├── src/
│   ├── api/                 # API client (snake_case → camelCase)
│   ├── components/          # UI, layout, feature components
│   ├── hooks/               # TanStack Query hooks
│   ├── pages/               # Route pages
│   └── types/               # TypeScript interfaces
├── Dockerfile               # Multi-stage (node → nginx)
└── nginx.conf               # SPA routing
```

## Docker

```bash
# Backend
docker build -f sportstream-api/deployments/Dockerfile sportstream-api/ -t sportstream-api
docker run -p 8080:8080 sportstream-api

# Frontend
docker build --build-arg VITE_API_URL=https://your-api-domain.com sportstream-dashboard/ -t sportstream-dashboard
docker run -p 80:80 sportstream-dashboard
```

## Environment Variables

### Backend
| Variable | Default | Description |
|----------|---------|-------------|
| `APP_PORT` | `8080` | Server port |
| `APP_ENV` | `development` | Environment |
| `APP_VERSION` | `1.0.0` | API version |
| `CORS_ALLOWED_ORIGINS` | `*` | Allowed origins (comma-separated) |

### Frontend
| Variable | Default | Description |
|----------|---------|-------------|
| `VITE_API_URL` | `http://localhost:8080` | Backend API URL (baked at build time) |

## Design

UI designed with [Google Stitch](https://stitch.withgoogle.com/) following patterns from Sportradar, SofaScore, and DAZN. Light theme with Material Design 3 color system.

Design files: `doc/design/dashboard/`

## License

MIT
