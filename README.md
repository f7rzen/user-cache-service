# User Cache Service

User Cache Service — REST-сервис на Go, который получает данные пользователей из внешнего API, кэширует их в памяти и предоставляет собственное API для получения списка пользователей и пользователя по ID.

## Реализовано

| Требование | Статус | Реализация |
|---|---|---|
| Получение списка пользователей | ✅ | GET /api/v1/users |
| Получение пользователя по ID | ✅ | GET /api/v1/users/{id} |
| Собственная модель ответа | ✅ | Ответ внешнего API преобразуется в UserResponse |
| In-Memory Cache | ✅ | Кэш на map |
| TTL | ✅ | Записи кэшируются на 1 минуту |
| Потокобезопасность | ✅ | Используется sync.RWMutex |
| Таймауты внешнего API | ✅ | Используется HTTP Client Timeout |
| Логирование HTTP-запросов | ✅ | Gin Logger |
| Graceful Shutdown | ✅ | Корректное завершение приложения |
| Docker | ✅ | Dockerfile |
| Docker Compose | ✅ | docker-compose.yml |
| Unit Tests | ✅ | Тесты основных сценариев Cache |

## Стек

- Go
- Gin
- In-Memory Cache
- sync.RWMutex
- slog
- Docker
- Docker Compose

## Структура проекта

```text
.
├── Dockerfile
├── README.md
├── cmd
│   └── app
│       └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── user-cache-service.postman_collection.json
└── internal
    ├── cache
    │   ├── cache.go
    │   └── cache_test.go
    ├── client
    │   └── user_client.go
    ├── config
    │   └── config.go
    ├── handler
    │   └── user_handler.go
    ├── model
    │   └── user.go
    └── service
        └── user_service.go
```

## Быстрый старт

Скачать проект:

```bash
git clone https://github.com/f7rzen/user-cache-service.git
cd user-cache-service
```

Запустить сервис:

```bash
docker compose up --build
```

После запуска API будет доступно по адресу:

```text
http://localhost:8080
```

Остановить сервис:

```bash
docker compose down
```

## Конфигурация

Сервис поддерживает конфигурацию через переменные окружения.

Пример:

```env
APP_PORT=8080
EXTERNAL_API_URL=https://jsonplaceholder.typicode.com
HTTP_CLIENT_TIMEOUT=5s
CACHE_TTL=1m
```

## API

### Проверка состояния сервиса

```http
GET /health
```

Пример curl:

```bash
curl http://localhost:8080/health
```

Пример ответа:

```json
{
  "status": "ok"
}
```

---

### Получение списка пользователей

```http
GET /api/v1/users
```

Пример curl:

```bash
curl http://localhost:8080/api/v1/users
```

Пример ответа:

```json
[
  {
    "id": 1,
    "full_name": "Leanne Graham",
    "email": "Sincere@april.biz",
    "city": "Gwenborough",
    "company": "Romaguera-Crona"
  },
  {
    "id": 2,
    "full_name": "Ervin Howell",
    "email": "Shanna@melissa.tv",
    "city": "Wisokyburgh",
    "company": "Deckow-Crist"
  },
  {
    "id": 3,
    "full_name": "Clementine Bauch",
    "email": "Nathan@yesenia.net",
    "city": "McKenziehaven",
    "company": "Romaguera-Jacobson"
  },
  {
    "id": 4,
    "full_name": "Patricia Lebsack",
    "email": "Julianne.OConner@kory.org",
    "city": "South Elvis",
    "company": "Robel-Corkery"
  },
  {
    "id": 5,
    "full_name": "Chelsey Dietrich",
    "email": "Lucio_Hettinger@annie.ca",
    "city": "Roscoeview",
    "company": "Keebler LLC"
  }
]
```

---

### Получение пользователя по ID

```http
GET /api/v1/users/{id}
```

Пример curl:

```bash
curl http://localhost:8080/api/v1/users/1
```

Пример ответа:

```json
{
  "id": 1,
  "full_name": "Leanne Graham",
  "email": "Sincere@april.biz",
  "city": "Gwenborough",
  "company": "Romaguera-Crona"
}
```

## Запуск тестов

```bash
go test ./...
```

В проекте реализованы unit-тесты для кэша:

- сохранение и получение значения;
- получение отсутствующего ключа;
- проверка истечения TTL.

## Postman collection

В корне проекта находится файл:

```text
user-cache-service.postman_collection.json
```

Чтобы использовать коллекцию:

1. Открыть Postman.
2. Нажать `Import`.
3. Выбрать файл `user-cache-service.postman_collection.json`.
4. Запустить нужный запрос.