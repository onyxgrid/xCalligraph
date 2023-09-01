run:
	docker-compose down && docker-compose build && docker-compose up -d

test:
	docker-compose down && go run cmd/main.go test

stop:
	docker-compose down