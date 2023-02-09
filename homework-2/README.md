## Homework 2

### Установка `ingress-nginx`
```shell
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace
```

### Применение
```shell
kubectl apply -f .
```
В `etc/hosts` необходимо прописать:
```shell
127.0.0.1 arch.homework
```

`GET http://arch.homework/health`
```json
{
    "status": "OK"
}
```
`GET http://arch.homework/otusapp/chelout/health`
```json
{
    "status": "OK"
}
```

[Коллекция Postman с тестами](tests.postman_collection.json)