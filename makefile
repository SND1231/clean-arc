.PHONY: up
up:
	docker-compose up -d --remove-orphans

.PHONY: stop
stop:
	docker-compose stop
