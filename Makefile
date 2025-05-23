include .env

BINARY_NAME=quotes-service
DOCKERFILE_PATH=build/docker/Dockerfile
COMPOSE_FILE=./docker-compose.yml

# Сборка Go-приложения
build:
	@echo "🔨 Сборка приложения..."
	@go build -o $(BINARY_NAME) ./cmd/main.go
	@echo "✅ Сборка завершена: ./$(BINARY_NAME)"

# Очистка бинарника
clean:
	@rm -f $(BINARY_NAME)
	@echo "🧹 Удален бинарник $(BINARY_NAME)"

# Очистка и запуск Go-кода
clear-start:
	@clear && go run ./cmd/main.go

# Сборка docker-образа
docker-build:
	@docker build -f $(DOCKERFILE_PATH) -t $(BINARY_NAME):latest .

# Запуск контейнеров
start:
	@docker compose --env-file ./.env -f $(COMPOSE_FILE) up -d

# Остановка контейнеров
stop:
	@docker compose --env-file ./.env -f $(COMPOSE_FILE) down

# Рестарт контейнеров
restart:
	@docker compose --env-file ./.env -f $(COMPOSE_FILE) down
	@docker rmi $(BINARY_NAME):latest || true
	@docker compose --env-file ./.env -f $(COMPOSE_FILE) up -d

# Просмотр логов
logs:
	@docker compose --env-file ./.env -f $(COMPOSE_FILE) logs -f app
