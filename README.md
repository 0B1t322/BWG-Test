# Как запустить
В корне прописать команду
```bash
docker-compose up --build
```

# Конфигурация
```yml
...
back:
    build: ./src
    environment:
        # порт для http приложения
      - BWG_APP_PORT=8080
        # строка подключения к бд
      - BWG_APP_POSTGRESQL_DSN=postgresql://user:password@db:5432/db?sslmode=disable
    ports:
      - 8080:8080
    depends_on:
      - db
    restart: always
...
```

# Swagger
После запуска приложения документация swagger доступна по пути:
http://localhost:8080/api/swagger

# Пайплайн для проверки работаспособности
1. Создать пользователя
2. Добавить ему сколько угодно разных транзакций
3. Получить транзакции для пользователя чтобы посмотреть результат по ним
4. Получить информацию о пользователе чтобы посмотреть его баланс.