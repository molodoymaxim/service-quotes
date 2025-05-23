# Импорт переменных окружения из .env
include .env
export $(shell sed 's/=.*//' .env)

# Очистка терминала и запуск локально
clear-start:
	@clear && go run ./cmd

# Запуск всех контейнеров в фоне
start:
	@docker compose --env-file .env -f docker-compose.yml up -d

# Остановка и удаление контейнеров
stop:
	@docker compose --env-file .env -f docker-compose.yml down

# Перезапуск приложения с удалением образа сервиса
restart:
	@docker compose --env-file .env -f docker-compose.yml down
	@docker rmi service_quote:latest
	@docker compose --env-file .env -f docker-compose.yml up -d

# Выполнение миграций
migrate:
	@docker compose --env-file .env -f docker-compose.yml run --rm migrate

createdb:
	docker exec -it postgres psql -U postgres -c "CREATE DATABASE quotes;"

startm:
	make createdb
	make migrate
	make start

# Логи сервиса
logs:
	@docker compose --env-file .env -f docker-compose.yml logs -f service_quote
