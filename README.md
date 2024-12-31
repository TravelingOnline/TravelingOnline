# Traveling Online System

This project is a Golang-based online traveling system that uses Docker for containerization. The application connects to PostgreSQL database on port 8080.

For installation follow the instructions below:

### 1.Clone the Repository

```bash
git clone <repository-url>
cd <repository-name>
```

### 3. set the env variables in .env file

```bash
POSTGRES_HOST=localhost
POSTGRES_USER=you-username
POSTGRES_PASSWORD=your-password
POSTGRES_DB_NAME=your-db-database
POSTGRES_PORT=5432
CONFIG_PATH=./config.json
```

### 2.Run the program

```bash
docker compose up -d --build
```
