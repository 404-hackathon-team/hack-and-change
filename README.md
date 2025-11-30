# Репозиторий команды "404"

Учебная платформа для онлайн обучения с возможностью загрузки домашних заданий, оценивания и уведомлениями. Трек "Студент".


## Как запустить:
1. Скопировать пример окружения (.env.template) и заполнить им файл .env:

2. Собрать и поднять все контейнеры из корня проекта:

Учтите, что вам нужно иметь установленный [Docker](https://www.docker.com)
```bash
docker compose up --build
```

3. Проверка работоспособности:
    `http://localhost:3000`


## Основные технологии:

**Backend**: `go`, `gin`, `goose`, `PostgreSQL`, `jwt`

**Frontend**: `Nuxt 4`, `Vue 3`, `TypeScript` 

---

## Frontend

[Ссылка на Figma](https://www.figma.com/design/zu5ewdmESV3CU5D3vX1uDg/hackaton-2-hack-change25?node-id=5-140&p=f&t=4oCaC9mmJaqvbuj5-0)

![photo_2025-11-30_05-32-09](https://github.com/user-attachments/assets/3f8818ee-9ef7-446c-9afa-5c99a3a99fbb)


### Структура проекта:
- `app/`>
- `components/` - Vue компоненты, используемые в приложении.
- `pages/` - Страницы приложения, которые автоматически маршрутизируются.
- `assets/` - Статические файлы.
- `composables/` - Повторно используемые функции и логика.
- `layouts/` - Макеты страниц.
- `app.vue` - Главный компонент приложения.

## Используемые технологии:
- [Nuxt 4](https://nuxt.com/) - Фреймворк
- [Vue 3](https://vuejs.org/) - Фреймворк
- [TypeScript](https://www.typescriptlang.org/) - Язык программирования
- [UnoCSS](https://unocss.dev/) - Утилитарный CSS фреймворк

---

## Backend

### Структура проекта:
- `/cmd/api` - инициализация REST API
- `/cmd/migrations` - миграции для разворачивания БД из контейнера
- `/config` - создание конфигурации программы на основе переменных окружения
- `/db` - инициализация подключение к базе данных
- `/service/auth` - работа с jwt-токенами и хэширование паролей
- `/service/course` - отправка запросов к базе данных и настрока Handler
- `/notification` - веб-сокет для уведомлений
- `/user` - обработка HTTP запросов
- `/types` - структура уведомлений и сущности БД
- `/utils/utils.go` - вспомогательные функции для сериализации данных в формат JSON и отправки ответов/ошибок на запрос
- `/utils/dataprovider.go` - структура для манипуляций с различного файлами на стороне сервера, вроде получения их от студента, регистрация файлов в базе данных, отправка данных пользователю и т.д.



## Используемые технологии:
- [gin](https://gin-gonic.com) - Фреймворк для работы с HTTP
- [goose](https://github.com/pressly/goose) - Библиотека для миграции бд
- [pgx](https://github.com/jackc/pgx) - Драйвер для работы с PostgreSQL
- [jwt](https://github.com/golang-jwt/jwt) - Библиотека создания jwt токена
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Библиотека хеширования
- [PostgreSQL](https://www.postgresql.org) - База данных PostgreSQL


---

## Состав команды "404":
1. Миронов Борис - Капитан, Backend разработчик, Devops [@Jeno7u](https://github.com/Jeno7u)
2. Михаил Бунто - Backend разработчик, Database Designer [@Undeadguy0](https://github.com/Undeadguy0)
3. Суслов Максим - Frontend раработчик, UI/UX [@ms0ur](https://github.com/ms0ur)

