# Go Post Flow

## Описание

Система для добавления и чтения постов и комментариев с использованием GraphQL.
Реализует функциональность, аналогичную комментариям к постам на популярных платформах (Хабр, Reddit).
Поддерживает иерархические комментарии, возможность отключения комментариев автором поста и асинхронную доставку новых комментариев через GraphQL Subscriptions.

---

## Особенности

- Просмотр списка постов.
- Просмотр поста и его комментариев.
- Возможность запрета комментариев.
- Комментарии организованы иерархически.
- Ограничение длины комментария (до 2000 символов).
- Пагинация для списка постов и комментариев.
- Поддержка GraphQL Subscriptions для асинхронной доставки комментариев.
- Выбор хранилища (in-memory или PostgreSQL).
- Использование Docker и docker-compose для развертывания.
- Покрытие функционала unit-тестами.
- Оптимизированная работа с вложенными комментариями и минимизация n+1 запросов используя dataloader.

---

## Быстрый старт

1. **Клонирование репозитория**:
    ```sh
    git clone https://github.com/Sergey-Polishchenko/go-post-flow
    cd go-post-flow
    ```
2. **Конфигурация**:
    ```sh
    cp .env.example .env
    ```
    + подробнее в [конфигурации](#конфигурация)
3. **Установка пакетов**:
    ```sh
    go mod tidy
    ```
4. **Генерация кода для GraphQL**:
    ```sh
    go run github.com/99designs/gqlgen generate
    ```
5. **Запуск**:
    ```sh
    docker build -t app .
    # запуск вручную:
    # --inmemory для запуска в режиме in-memory без данного флага будет использоваться конфигурация удаленной базы данных
    docker run -p 8080:8080 app --inmemory
    # запуск с использованием docker-compose:
    # только в режиме PostgreSQL c инициализацией контейнера базы данных
    docker-compose up --build
    ```
    + подробнее в [развертывании](#развертывание)
6. **Подключение**:
    Для простоты playground: [localhost:8080](http://localhost:8080/) // Если вы изменили конфиг укажите свой порт

    Примеры запросов GraphQL находятся в [директории с примерными запросами](./graphql-test-queries/)

    + подробнее в [запросах](#запросы)

---

## Зависимости

- [Go 1.23.4](https://go.dev)
- [gqlgen](https://github.com/99designs/gqlgen)
- [Docker](https://www.docker.com)
- [task](https://github.com/go-task/task) (опционально для тестирования и развертывания)

---

## Конфигурация

### Переменные окружения

**Application**:
PORT=8080 - порт на котором работает приложение

**PostgreSQL** - настройки базы данных
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=mydb
DB_PORT=5432
DB_HOST=postgres_db

**Можете просто скопировать переменные окружения из примера**:
```sh
cp .env.example .env
```

---

## Тестирование

**Для тестирования**:
```sh
go test ./...
# если используете task
task test
```

---

## Развертывание

### С использованием Task

```sh
# сборка Docker образа
task build
# запуск в режиме PostgreSQL
task run
# запуск в режиме in-memory
task run-inmemroy

# сборка и развертывание docker-compose вместе с миграцией и запуском PostgreSQL базы данных
task up
# очистка сборки docker-compose
task down
```

### Без использования Task

```sh
# сборка Docker образа
docker build -t myapp .
# запуск в режиме PostgreSQL
docker run -p 8080:8080 myapp
# запуск в режиме in-memory
docker run -p 8080:8080 myapp --inmemory

# сборка и развертывание docker-compose вместе с миграцией и запуском PostgreSQL базы данных
docker-compose up --build
# очистка сборки docker-compose
docker-compose down
```

---

## Запросы

**Запросы**:
- createPost(input: PostInput!) => Post!
- createComment(input: CommentInput!) => Comment!
- createUser(input: UserInput!) => User!
- post(id: ID!) => Post!
- posts(limit: Int, offset: Int) => [Post!]!
- comment(id: ID!) => Comment!
- user(id: ID!) => User!
**Подписка**:
- commentAdded(postId: ID!) => Comment!

Примеры запросов GraphQL находятся в [директории с примерными запросами](./graphql-test-queries/)

Так же можно посмотреть [схему](./schema/)

### Примеры запросов через curl

**Создание нового пользователя**
```sh
curl -X POST http://localhost:8080/query \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation CreateUser { createUser(input: { name: \"name\" }) { id name } }"
  }'
```

**Создание нового поста**
```sh
curl -X POST http://localhost:8080/query \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation CreatePost { createPost(input: { title: \"title\", content: \"content\", authorId: \"1\", allowComments: true }) { id title } }"
  }'
```

**Создание нового комментария**
```sh
curl -X POST http://localhost:8080/query \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation CreateComment { createComment(input: { text: \"text\", postId: \"1\", authorId: \"1\" }) { id text } }"
  }'
```

**Получение постов**
```sh
curl -X POST http://localhost:8080/query \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query GetPost { posts { id title content allowComments author { id name } comments { id text } } }"
  }'
```

- **Подписка на пост требует WebSocker поэтому через curl ее оформить не удастся**

---

## Структура проекта

```sh
├── README.md                   # Документация проекта
├── LICENSE                     # Лицензия проекта
├── go.mod                      # Go модуль для зависимостей
├── go.sum                      # Контрольные суммы зависимостей
├── Dockerfile                  # Конфигурация Docker для создания контейнера
├── docker-compose.yml          # Конфигурация Docker Compose для запуска проекта
├── Taskfile.yml                # Конфигурация Task для автоматизации задач
├── graphql-test-queries/       # Примеры GraphQL запросов для тестирования
├── gqlgen.yml                  # Конфигурация для генерации GraphQL схемы
├── schema/                     # GraphQL схемы
├── migrations/                 # Миграции PostgreSQL
├── cmd/                        # Основной код приложения
│   └── server/                 # Логика сервера
│       └── server.go           # Входная точка проекта
└── internal/                   # Внутренняя логика приложения
    ├── broadcast/              # Логика для рассылки уведомлений
    ├── config/                 # Конфигурация и обработка флагов
    ├── delivery/graph/         # GraphQL модели, резолверы и даталоадеры
    ├── errors/                 # Централизированные ошибки
    ├── repository/             # Хранилища данных (in-memory и PostgreSQL)
    ├── tools/                  # Пакеты для генерации gqlgen
    └── utils/                  # Утилита пагинации
```

---

## TODO

- [ ] увеличить покрытие тестами
- [ ] протестировать на большую нагрузку

---

## Лицензия

Этот проект распространяется под лицензией [MIT](LICENSE).
