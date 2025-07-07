# MoneyFlow Go Backend

## Запуск

1. Скопируйте пример файла переменных окружения и настройте значения:

```bash
cp .env.example .env
```

2. Отредактируйте файл `.env` при необходимости (например, смените JWT_SECRET или параметры БД).

3. Соберите и запустите сервисы:

```bash
docker-compose up --build
```

4. Примените миграции (внутри контейнера backend):

```bash
docker-compose exec backend goose -dir ./migrations postgres "host=$DB_HOST user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable" up
```

## Необходимые переменные окружения

- `DB_HOST` — адрес сервера БД (по умолчанию: db)
- `DB_PORT` — порт БД (по умолчанию: 5432)
- `DB_USER` — пользователь БД
- `DB_PASSWORD` — пароль пользователя БД
- `DB_NAME` — имя базы данных
- `DB_URL` — строка подключения к БД (например: postgres://moneyflow_user:moneyflow_pass@db:5432/moneyflow?sslmode=disable)
- `JWT_SECRET` — секрет для подписи JWT (обязательно смените в продакшене!)

## Эндпоинты

- `POST /register` — регистрация пользователя
- `POST /login` — логин (возвращает access и refresh токены)
- `POST /logout` — логаут (требует refresh_token в теле запроса)
- `GET /health-check` — проверка статуса сервиса (не входит в Swagger)

### Пример запроса на логаут

```json
{
  "refresh_token": "<refresh_token>"
}
```

## Swagger

Swagger-документация доступна по адресу: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
