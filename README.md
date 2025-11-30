# Репозиторий команды "404"

---

## Frontend

Информация о фронтенд части приложения [тут(клик)](https://github.com/404-hackathon-team/hack-and-change/blob/master/frontend/README.md)

[Ссылка на Figma](https://www.figma.com/design/zu5ewdmESV3CU5D3vX1uDg/hackaton-2-hack-change25?node-id=5-140&p=f&t=4oCaC9mmJaqvbuj5-0)

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
---


## Стек backend

Язык программировния: GO

БД: PostgresSQL

---

## Структура

/cmd/api - инициализация REST API

/cmd/migrations - миграции для разворачивания БД из контейнера

/config - создание конфигурации программы на основе переменных окружения

/db - инициализация подключение к базе данных

/service/auth - работа с jwt-токенами и хэширование паролей

/service/course - отправка запросов к базе данных и настрока Handler

/notification - веб-сокет для уведомлений

/user - обработка HTTP запросов

/types - структура уведомлений и сущности БД

/utils/utils.go - вспомогательные функции для сериализации данных в формат JSON и отправки ответов/ошибок на запрос

/utils/dataprovider.go - структура для манипуляций с различного файлами на стороне сервера, вроде получения их от студента, регистрация файлов в базе данных, отправка данных пользователю и т.д.


## Состав команды
1. Миронов Борис - Капитан, Backend разработчик, Devops
2. Михаил Бунто - Backend разработчик
3. Суслов Максим - Frontend раработчик, UI/UX

