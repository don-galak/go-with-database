steps

1. create Dockerfile, docker-compose.yml, main.go
2. docker compose build
3. docker compose run -rm app go mod init github.com/${user}/${project-name}
4. docker compose run --rm app air init

Run with: `docker compose up`

Update imports: `docker compose run --rm app go mod tidy`

Create migration: `docker compose --profile tools run create-migration create_items`

Run migration: `docker compose --profile tools run migrate`

Access postgres: `docker compose exec db psql -U local-dev -d api`

then, retrieve table: `\d ${tablename};`

Exit postgres: `\q`

Create item: `curl -F "name=${name}" http://localhost:3000/items`

Retrieve items: `curl http://localhost:3000/items`

link: [Develop a Go app with Docker Compose](https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/)