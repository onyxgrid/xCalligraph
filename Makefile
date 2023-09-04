run:
	docker-compose down && docker-compose build && docker-compose up -d

test:
	docker-compose down && go run cmd/main.go test

testwithsigning:
	docker-compose down && go run cmd/main.go testwithsigning

stop:
	docker-compose down