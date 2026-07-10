# Анскилл API

Чисто потыкать фрейморк и реализовать basic CRUDs

## Стек

- **Go** 1.25
- **Fiber v2** — веб-фреймворк
- **GORM** — ORM
- **SQLite** — база данных

## Структура проекта

```
simple_api/
├── cmd/
│   └── main.go                  # Точка входа
├── app/
│   ├── app.go                   # Инициализация приложения и роутов
│   ├── api/
│   │   └── books/
│   │       ├── handlers.go      # HTTP-обработчики
│   │       ├── routers.go       # Регистрация роутов
│   │       └── response.go      # Структуры ответов
│   ├── dto/
│   │   └── book.go              # DTO для запросов
│   ├── service/
│   │   └── books.go             # Бизнес-логика
│   ├── repository/
│   │   └── books.go             # Работа с БД
│   ├── model/
│   │   └── book.go              # GORM-модель
│   ├── config/
│   │   └── config.go            # Конфигурация
│   └── util/
│       └── validator/
│           └── validator.go     # Валидация
└── pkg/
    └── database/
        └── database.go          # Подключение к БД
```

## Запуск

```bash
go run cmd/main.go
```

Сервер запускается на порту `:3000`.

## Конфигурация

Настройки задаются в файле `.env`:

```env
VERSION=v1
HOST=localhost
PORT=8000
```

## API

Base URL: `http://localhost:3000/api/v1`

### Книги

| Метод    | Путь          | Описание             |
|----------|---------------|----------------------|
| `GET`    | `/book/:id`   | Получить книгу по ID |
| `POST`   | `/book`       | Создать книгу        |
| `PUT`    | `/book/:id`   | Обновить книгу       |
| `DELETE` | `/book/:id`   | Удалить книгу        |

### Примеры запросов

**Создать книгу**
```http
POST /api/v1/book
Content-Type: application/json

{
  "title": "Война и мир",
  "author": "Лев Толстой",
  "rating": 5
}
```

**Получить книгу**
```http
GET /api/v1/book/1
```

**Обновить книгу**
```http
PUT /api/v1/book/1
Content-Type: application/json

{
  "title": "Новое название",
  "author": "Автор",
  "rating": 4
}
```

**Удалить книгу**
```http
DELETE /api/v1/book/1
```

### Модель книги

| Поле      | Тип    | Описание        |
|-----------|--------|-----------------|
| `id`      | int    | Идентификатор   |
| `title`   | string | Название        |
| `author`  | string | Автор           |
| `rating`  | int    | Рейтинг         |
