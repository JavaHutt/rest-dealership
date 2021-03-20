up	:
	docker-compose up -d

down	:
	docker-compose down

run	:
	go run cmd/main.go

seed	:
	http POST http://localhost:8000/db/seed < ./internal/data/mocks/vehicles.json
