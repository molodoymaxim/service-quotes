# Запуск всех контейнеров в фоне (Postgres + migrate + сервис)
start:
	docker compose --env-file .env -f docker-compose.yml up -d --build

# Остановка и удаление контейнеров
stop:
	docker compose --env-file .env -f docker-compose.yml down

# Перезапуск приложения с удалением образа сервиса
restart: stop
	docker rmi service_quote:latest || true
	$(MAKE) start

# Логи сервиса
logs:
	docker compose --env-file .env -f docker-compose.yml logs -f service_quote
