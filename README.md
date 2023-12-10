steps

1. create Dockerfile, docker-compose.yml, main.go
2. docker compose build
3. docker compose run -rm app go mod init github.com/${user}/${project-name}
4. docker compose run --rm app air init

Run with: `docker compose up`

Run to update imports: `docker compose run --rm app go mod tidy`

link: [Develop a Go app with Docker Compose](https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/)