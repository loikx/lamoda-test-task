up:
	docker compose up -d
	sleep 5
	docker ps

down:
	docker compose down

test:
	go test ./...