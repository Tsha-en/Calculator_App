# Web - Calculator
Привет! Это проект веб-сервера, который вычисляет простые математические выражения.
Работает он следующим образом - на сервер отправляются GET-запросы с математическим
выражением, в ответ приходит http - ответ с телом в виде:
___{
    "result": "результат выражения"
}___
Или ошибка, если таковая случилась.

## Инструкция по запуску проекта.

1.  Скачайте все файлы проекта в какую нибудь папку на вашем компьютере.
2.  Откройте эту папку в вашем редакторе кода.
3.  В терминале выполните команду: 
```
go run cmd/main.go
```
4. Чтобы отправить запрос на сервер, можете воспользоваться следующим кодом,
просто введя его в терминал:

```java
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+(2*2)"
}'
```
Если не работает, попробуй:

```java
curl -X POST http://localhost/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1+1\"}"
```

Но самый лучший и удобыный вариант, воспользоваться программой Postman.

## Структура проекта

* cmd - тут точка входа в программу, то есть main.go
* internal/application - внутренняя логика программы(http, запуск, остановка и т.д.)
* pkg/calculator - тут файлы с кодом канкуляторя, и файл с тестами для него.
* примеры - в соответствии с условиями, прилагаются примеры со всеми возможными сценариями 
выполнения программы.
* go.mod
* README.md - вы сейчас тут)