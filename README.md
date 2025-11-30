# Репозиторий команды "404"

---

## Frontend

Информация о фронтенд части приложения [тут(клик)](https://github.com/404-hackathon-team/hack-and-change/blob/master/frontend/README.md)

---

## Backend

Создание контейнеров
```bash
docker-compose up --build
```

Запуск тестов
```bash
go test -v ./...
```

Миграция происходит при запуске сервера через goose

Установка goose
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Создать новый файл миграции 
```
goose create [migration_name] sql
```

