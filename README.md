## Necessary steps, execute in order:

1. create Dockerfile, docker-compose.yml, main.go
2. docker compose build
3. docker compose run -rm app go mod init github.com/${user}/${project-name}
4. docker compose run --rm app air init

## General commands

Run with: `docker compose up`

Update imports: `docker compose run --rm app go mod tidy`

Create migration: `docker compose --profile tools run create-migration create_items`

Run migration: `docker compose --profile tools run migrate`

Access postgres: `docker compose exec postgres psql -U local-dev -d api`

then, retrieve table: `\d ${tablename};`

Exit postgres: `\q`

Create item: `curl -F "name=${name}" http://localhost:3000/items`

Retrieve items: `curl http://localhost:3000/items`

Build: `docker build -t go-with-database .`

Run build: `docker run -e API_SERVER_ADDR=:3000 go-with-database`

## References

link: [Develop a Go app with Docker Compose](https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/)