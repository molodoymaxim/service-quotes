
# Данное приложение представляет собой REST API-сервис на Go для хранения и управления цитатами.

## Функциональные требования:
- Добавление новой цитаты (POST /quotes)
- Получение всех цитат (GET /quotes)
- Получение случайной цитаты (GET /quotes/random)
- Фильтрация по автору (GET /quotes?author=Confucius)
- Удаление цитаты по ID (DELETE /quotes/{id})

## Проверочные команды (curl):

Добавление новой цитаты
```
curl -X POST http://localhost:8080/quotes \
-H "Content-Type: application/json" \
-d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

Получение всех цитат
```
curl http://localhost:8080/quotes
```

Получение случайной цитаты
```
curl http://localhost:8080/quotes/random
```

Фильтрация по автору
```
curl http://localhost:8080/quotes?author=Confucius
```

Удаление цитаты по ID
```
curl -X DELETE http://localhost:8080/quotes/1
```
## Запуск
Для запуска необходимо клонировать репозиторий
```
https://github.com/molodoymaxim/service-quotes.git
```

Создать .env файл в корне проекта, содержащий переменны окружения.
Пример:
```
HTTP_PORT=8080
HTTP_TIME_LIFE_CTX=1

# Подключени к PostgreSQL
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=1234
POSTGRES_DB=quotes

# Настройка соединения с БД Postgres
DB_MAX_CONN_PS=100 
DB_CONN_IDLE_TIME_PS=60
DB_CONN_LIFE_TIME_PS=120
DB_QUERY_TIMEOUT_PS=60
```
Выполнить команду запуска проекта.
Запуск всех контейнеров в фоне (Postgres + migrate + сервис)
```
make start
```
Дополнительные команды:
Остановка и удаление контейнеров
```
make stop
```

Перезапуск приложения с удалением образа сервиса
```
make restart
```

Логи сервиса
```
make logs
```
