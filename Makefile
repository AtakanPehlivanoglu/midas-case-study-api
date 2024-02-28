dev-up:
	cd ./scripts/dev; \
	docker-compose up -d

dev-down:
	cd ./scripts/dev; \
	docker-compose down