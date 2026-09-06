# image-generator
A golang based image generator and image feature adding software specially to work with the image files.


## Backend Setup

### 1. Install Docker

Install Docker and Docker Compose on your system.

Verify the installation:

```bash
docker --version
docker compose version
```

### 2. Clone the Repository

```bash
git clone <your-repository-url>
cd image-generator/backend
```

### 3. Start PostgreSQL

Start the PostgreSQL 16 container:

```bash
docker compose up -d
```

Check that the container is running:

```bash
docker ps
```

You should see:

```text
backend-db-1   postgres:16   ...   0.0.0.0:5432->5432/tcp
```

### 4. Connect to PostgreSQL

Open the PostgreSQL shell inside the Docker container:

```bash
docker exec -it backend-db-1 psql -U myuser -d pixelForge
```

### 5. Run the Backend

From the backend directory:

```bash
go run .
```

### 6. Stop PostgreSQL

When finished:

```bash
docker compose down
```

