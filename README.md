# Adtelligent Test Task by Maksym Zhovtaniuk

### Техническое задание для кандидатов:

1. разработка хттп сервера (задание обязательно к выполнению)

Представьте, что вы получили новый проект интернет-магазина и вам нужно заложить архитектуру для его разработки и
поддержки. В качестве тестовго задания полностью спроектируйте базу данных, а так же сделайте CRUD одной (любой)
сущности. HTTP сервер должен быть написан на GoLang, масимально просто, без использования фреймворков.

Техническое Задание:

- разработать HTTP API с базовой авторизацией, которое будет позволять выполнять CRUD операции над сущностями.
  Пользователь будет один (администратор, который и будет создавать эти сущности)
- формат ответа: JSON
- описание сущностей и полей (если вы считаете, что какого-то поля не хватает, вы можете смело его добавить):
    - продавец (имя, телефон)
    - товар (название, описание, цена, продавец)
    - покупатель (имя, телефон)
    - заказ (покупатель, несколько товаров)


2. оптимизация функции конкатенации. (задание со звездочкой, можно не делать, если не знаете)

Опмтимизируйте скорость выполнения функции. Кол-во значений во входящем параметре (len(str)) >= 30.
Напишите бенчмарк тест на эту функцию и на её оптимизированную версию.

```
func concat(str []string) string  {
    result := ""
    for _, v := range str {
        result += v
    }
    return result
}
```

Выполненное тестовое задание разместите на гитхабе.
Доступ к проекту предоставьте на аккаунт: https://github.com/Kirill-Shkodkin

---

### Task 1. HTTP Server

Create .env file with the following values

```dotenv
HTTP_HOST=<host>

DB_USER=<db_user>
DB_PASSWORD=<db_password>

JWT_SIGNING_KEY=<signign_key>
PASSWORD_SALT=<password_salt>
```

### Steps to run the app

Run docker-compose file
```shell
docker-compose up -d --build app
```

Apply migrations to database
```shell
cd ./migrator
```

```shell
docker build -t app-migrator .
```

```shell
docker run --network host app-migrator -path=/schema -database "mysql://root:qwerty123@tcp(localhost:3306)/adtelligent-db" up
```

### Task 2 
### [Concat functions benchmark](task2)

| Slice len | Basic        | With Join  | With Strings Builder |
|-----------|--------------|------------|----------------------|
| 100       | 154416 ns/op | 3420 ns/op | 3252 ns/op           |
| 1000      | 13423015 ns/op            | 35641 ns/op          | 33679 ns/op                    |
| 10000     | 1581037667 ns/op            | 343412 ns/op          | 319548 ns/op                    |