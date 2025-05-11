# QuickDrop
QuickDrop is a lightweight file-sharing app that lets users upload a file and receive a unique download link. Designed to demonstrate file handling, storage, and link generation. Useful as a backend utility or quick sharing tool.

---

## Steps to run this locally:

1. Make sure you have golang installed with version > 1.22
2. Make sure you have npm and node installed
3. Download dependencies:
  - `cd backend && go mod tidy && cd ..`
  - `cd frontend && npm i && cd ..`
4. Update some variables [TODO: Use .env]
  - `docker-compose.yaml` : Update the env variables and volume location/name
  - `backend/internals/config/config.go`: Update config variables as required
  - `frontend/src/api/file.ts`: Update server url if required
5. Run the services:
  - `docker compose up -d`
  - `cd backend && go run cmd/api/main.go`
  - `cd frontend && npm run dev`