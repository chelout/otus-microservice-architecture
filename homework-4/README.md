## Homework 4

В данной домашней работе разворачивается Kubernetes кластер с `api`-приложением `users`.
В качестве зависимостей используются чарты `postgresql` и `ingress-nginx`. Инициализация бд происходит средствами чарта, а миграции вынесены в отдельный `job`. Особенность `job` заключается в использовании `init`-контейнера, который нужен для проверки соединения с бд. Как только бд готова к использованию, init-контейнер завершает свою работу и позволяет запуститься миграциям.
Для подов содержащих `api`-приложение `users` тоже используются `init`-контейнеры, они не дают запуститься подам пока не отработают миграции.

### Установка чарта с зависимостями
Установка чарта производится в отдельный неймспейс.
```shell
helm upgrade --install hw4-release . \
  --dependency-update \
  --create-namespace \
  --namespace=hw4 
```

В `etc/hosts` необходимо прописать:
```shell
127.0.0.1 arch.homework
```

Список доступных методов `api`:
- POST http://arch.homework/user
- GET http://arch.homework/user/{userId}
- PUT http://arch.homework/user/{userId}
- DELETE http://arch.homework/user/{userId}

Вспомогательные методы `api`:
- GET http://arch.homework/health
- GET http://arch.homework/hostname

[Коллекция Postman с тестами](Homework%204.postman_collection.json)