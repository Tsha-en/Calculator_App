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
2.  Откройте эту папку в терминале, или редакторе кода.
3.  В терминале выполните команду: 
```
go run cmd/main.go
```
или запустите файл main.go с помощью редактора кода.
Также, если у вас Linux/MacOS, можно просто запустить исполняемый файл main.

4. Чтобы отправить запрос на сервер, можете воспользоваться следующим кодом,
просто введя его в терминал:

```java
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+(2*2)"
}'
```
Если не работает:

```java
curl -X POST http://localhost/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1+1\"}"
```

Но самый лучший и удобыный вариант, воспользоваться программой Postman.

## Структура проекта

* cmd - тут точка входа в программу, то есть main.go
* internal/application - внутренняя логика программы(http, запуск, остановка и т.д.)
* pkg/calculator - тут файлы с кодом канкуляторя, и файл с тестами для него.
* go.mod
* README.md - вы сейчас тут)

## Примеры с ошибками

### Ошибка 500 
```java
curl -X POST http://localhost/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1+1"
```   


```java {
    "error": "Internal server error"
}```

### Ошибка 422 

 ```java 
 curl -X POST http://localhost/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"sd1\"}"
 ```


 ```java
{
    "error": "Expression is not valid"
}%  
```

### Нормальное выполнение программы

```java
curl -X POST http://localhost/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1+1\"}"
```
```java
{
        "result": "2"
}%                                   
```
По всем вопросам - @Antonenot