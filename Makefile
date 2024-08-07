run-app:
	docker compose -f deployments/docker-compose.yaml up -d;

build-app:
	go build -o app ./cmd/main.go && docker compose -f deployments/docker-compose.yaml build;