# Репозиторий команды "404"

---

# Backend

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

