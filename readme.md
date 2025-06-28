# One-Time Note

A simple one-time note service written in Go

## Features

- One-time storage with expiration
- Redis-based backend
- Dockerized and ready for deployment
- Environment-based configuration with optional YAML config

## Getting Started

1. Clone the repository:

    ```bash
   git clone https://github.com/radahn42/onetime-note.git
   cd onetime-note
   ```

2. Create .env from the example:

    ```bash
   cp .env.example .env
    ```

3. Adjust configuration if needed.
4. Run the app:

    ```bash
   docker-compose up --build
    ```

## Configuration

- `.env`: environment-specific variables (e.g., redis password)
- `config/config.yaml`: optional application config (overridden by env)

Example `config.yaml`:

   ```yaml
   app:
     addr: ":8080"
     
   redis:
     addr: "redis:6379"
     password: ""
   ```

## API Usage

### Create a note

   ```bash
   curl -X POST http://localhost:8080/api/notes \
     -H "Content-Type: application/json" \
     -d '{"content": "my secret message", "ttl_seconds": 60}'
   ```

**Response:**

   ```json
   {
      "id": "a1b2c3d4"
   }
   ```

### Get and destroy the note

   ```bash
   curl http://localhost:8080/api/notes
   ```

**Response:**

   ```json
   {
      "content": "my secret message"
   }
   ```
