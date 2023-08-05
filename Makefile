run:
	docker-compose down && docker-compose build && docker-compose up -d

stop:
	docker-compose down