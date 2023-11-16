# Тестовое задание One Day Offer МойОфис

Необходимо реализовать CLI-утилиту, которая реализует асинхронную обработку входящих URL из файла, переданного в качестве аргумента данной утилите.
Формат входного файла: на каждой строке – один URL. URL может быть очень много! Но могут быть и невалидные URL.

Пример входного файла:  
https://myoffice.ru  
https://yandex.ru  

По каждому URL получить контент и вывести в консоль его размер и время обработки. Предусмотреть обработку ошибок.

# Пример использования 

```bash
go run cmd/main.go
```

```bash
go run cmd/main.go -filepath ./testdata/urls.txt
```

Результат выполнения:
```bash
2023/11/17 23:44:22 ERROR cant verify URL error="not an URL" url=httpz://yandex.ru
2023/11/17 23:44:22 ERROR cant verify URL error="not an URL" url=//yandex.ru
2023/11/17 23:44:22 ERROR cant verify URL error="not valid scheme" url=yandex.ru
2023/11/17 23:44:22 ERROR cant do GET request error="Get \"http://localhost:8080\": dial tcp 127.0.0.1:8080: connect: connection refused"
2023/11/17 23:44:22 INFO length: 15671, duration: 169ms
```