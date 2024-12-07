# Шаги для запуска приложения

## Клонирование репозитория:

- git clone 

## Запуск приложения с помощью Docker Compose:

- docker-compose up --build



Отправьте POST запрос на /auth/token для генерации пары Access и Refresh токенов:
- curl -X POST http://localhost:8080/auth/token -d '{"user_id": "123e4567-e89b-12d3-a456-426614174000"}'

Отправьте POST запрос на /auth/refresh для обновления пары Access и Refresh токенов:
- curl -X POST http://localhost:8080/auth/refresh -d '{"refresh_token": "your_refresh_token"}'