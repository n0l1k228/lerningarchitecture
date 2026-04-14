include .env
export

migrate-up:
	@make migrate-action action=up
migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "action необходимый параметр"; \
		exit 1; \
	fi; \
		migrate \
		-path ./migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable \
		$(action)