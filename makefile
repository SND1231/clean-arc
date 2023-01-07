.PHONY: up
up:
	docker-compose up -d --remove-orphans

.PHONY: down
down:
	docker-compose down --volumes --remove-orphans
