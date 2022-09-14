.PHONY: dev dbup dbdown

dev:
	go run cmd/api/main.go

dbup:
	docker compose up -d postgres

dbdown:
	docker compose down postgres