# сокращатель ссылок
Тестовое задание OZON

<!-- ToC start -->
# Содержание

1. [Описание задачи](#Задание)
1. [Демонстрация работы](#Демонстрация работы)
3. [Сборка и запуск](#Сборка-и-запуск)
4. [Тестирование](#Тестирование)
<!-- ToC end -->



# Задание

Необходимо реализовать сервис, который должен предоставлять API по созданию сокращенных ссылок следующего формата:
- Ссылка должна быть уникальной и на один оригинальный URL должна ссылаться только одна сокращенная ссылка.
- Ссылка должна быть длинной 10 символов
- Ссылка должна состоять из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание)


Сервис должен быть написан на Go и принимать следующие запросы по http:
1. Метод Post, который будет сохранять оригинальный URL в базе и возвращать сокращённый
2. Метод Get, который будет принимать сокращённый URL и возвращать оригинальный URL

Решение должно быть предоставлено в «конечном виде», а именно:
- Сервис должен быть распространён в виде Docker-образа
- В качестве хранилища ожидается использовать in-memory решение И postgresql. Какое хранилище использовать указывается параметром при запуске сервиса.
- Покрыть реализованный функционал Unit-тестами

# Демонстрация работы 
![ozon_test](https://user-images.githubusercontent.com/91884862/182147188-d981301f-714a-422a-8d09-7d3445c6d769.gif)

**Структура проекта:**
```
.
├── internal
│   ├── database           // работа с postgres (загрузка/выгрузка в бд)
│   ├── handlers         	// обработка запросов (как postgres так и internal memory)
│   ├── models           	// структура JSON
│   ├── utils
│       ├── hash_func    	// хэш функция
│       ├── json_parse    	// парсинг JSON
│       ├── responses    	// обработчик ответов
├── postgres            	// SQL файлы с миграциями
└── configs             	// конфиг файл postgres
```

**Пример**

Запрос:

```
curl --request POST --data '{"url" : "ozon"}' http://localhost:8080/
```

Ответ:

```
{
   {"url":"Li0QUvKTcT"}

}
```


**Пример**

Запрос:

```
curl --request GET --data '{"url" : "Li0QUvKTcT"}' http://localhost:8080/
```

Ответ:

```
{
   {"url":"ozon"}
}
```
