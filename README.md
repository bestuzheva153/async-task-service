# Async Task Service

Асинхронный backend-сервис на Go для создания и обработки задач с использованием PostgreSQL и фонового worker'а.

## Технологии

- Go
- Gin
- PostgreSQL
- Docker
- Docker Compose

## Архитектура

Проект разделён на несколько слоёв:

- `handler` — обработка HTTP-запросов
- `service` — бизнес-логика
- `repository` — работа с базой данных
- `worker` — фоновая обработка задач
- `app` — инициализация зависимостей приложения
- `router` — регистрация маршрутов

## Структура проекта

```text
cmd/
  app/
    main.go

internal/
  app/
  config/
  http/
    handler/
    router/
  model/
  repository/
  service/
  worker/
