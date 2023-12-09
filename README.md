# Документация

### Запуск проекта
* Поднять docker
* Выполнить в файле в папке с проектом (сервер поднимиться по адресу: http://localhost:8081/, бд: postgres://admin:password@db:5432/product?sslmode=disable)
```shell
make up
```
* Чтобы остановить работу проекта 
```shell
make down
```

###  В ПРИМЕРАХ ЗАПРОСА УКАЗАН {ID} ЕГО НЕОБХОИМО ВЗЯТЬ ИЗ ОТВЕТОВ, КОТОРЫЕ ВАМ ПРИХОДЯТ

### Методы API
#### POST /api/warehouse/create RESPONSE 200

Тело запроса
```json5
{
  "name": "new warehouse name",
  "availability": true
}
```
Пример запроса
```shell
curl --location 'http://localhost:8081/api/warehouse/create' \
--header 'Content-Type: application/json' \
--data '{
    "name": "new warehouse name",
    "availability": true
}'
```

Пример ответа
```json5
{
    "data": {
        "type": "warehouse",
        "id": "018c50ba-0a30-7596-a8ae-7b0f9869581f",
        "attributes": {
            "name": "New warehouse",
            "availability": true
        }
    }
}
```
id - uuid тип (генерируется системой)

#### DELETE /api/warehouse/delete/{id} RESPONSE 204
id - uuid тип (уникальный номер склада, сгенерированный системой)

Пример запроса
```shell
curl --location --request DELETE 'http://localhost:8081/api/warehouse/delete/{id}'
```

В ОТВЕТЕ НИЧЕГО НЕ БУДЕТ

#### POST /api/products/create RESPONSE 200

Тело запроса
```json5
{
  "name": "string",
  "count": 1,
  "size": {
    "length": 1.25,
    "width": 2.0,
    "height": 3.15,
    "unit": "sm"
  },
  "warehouse": "uuid"
}
```

Пример запроса
```shell
curl --location 'http://localhost:8081/api/products/create' \
--header 'Content-Type: application/json' \
--data '{
    "name": "New product name",
    "count": 100,
    "size": {
        "length": 1.24,
        "width": 1,
        "height": 3.45,
        "unit": "sm"
    },
    "warehouse": "warehouse uuid id"
}'
```

Пример ответы
```json5
{
    "data": {
        "type": "product",
        "id": "018c50bd-3d16-7596-94ff-9ffb36798737",
        "attributes": {
            "name": "New product name",
            "count": 100,
            "size": {
                "length": 1.24,
                "width": 1,
                "height": 3.45,
                "unit": "sm"
            },
            "is_reserved": false
        },
        "relationships": {
            "warehouse": {
                "data": {
                    "type": "warehouse",
                    "id": "018c50ba-0a30-7596-a8ae-7b0f9869581f"
                }
            }
        }
    }
}
```
id - uuid тип (уникальный номер товара, сгенерированный системой)

#### DELETE /api/products/delete/{id} RESPONSE 204
id - uuid тип (уникальный номер товара, сгенерированный системой)

Пример запроса
```shell
curl --location --request DELETE 'http://localhost:8081/api/products/delete/{id}'
```

#### GET /api/products/find-by-warehouse/{id} RESPONSE 200
id - uuid тип (уникальный номер товара, сгенерированный системой)

Пример запроса
```shell
curl --location 'http://localhost:8081/api/products/find-by-warehouse/{id}'
```

Пример ответа (запрос валиден)
```json5
{
    "data": [
        {
            "type": "product",
            "id": "018c50cf-a9c1-75f1-acef-5c5381b77c80",
            "attributes": {
                "name": "New product name",
                "count": 100,
                "size": {
                    "length": 1.24,
                    "width": 1,
                    "height": 3.45,
                    "unit": "sm"
                },
                "is_reserved": false
            },
            "relationships": {
                "warehouse": {
                    "data": {
                        "type": "warehouse",
                        "id": "018c50cf-805c-75f1-afd6-11d2e0d0b3d7"
                    }
                }
            }
        }
    ],
    "count": 1
}
```

Пример ответа (запрос с ошибкой)
```json5
{
    "error": "uuid: incorrect UUID length 1 in string 1"
}
```

#### PATCH /api/products/reserve RESPONSE 200

Пример запроса
```shell
curl --location --request PATCH 'http://localhost:8081/api/products/reserve' \
--header 'Content-Type: application/json' \
--data '{
    "ids": [
        "product id"
    ]
}'
```

В ОТВЕТЕ НИЧЕГО НЕ БУДЕТ

#### PATH /api/products/release RESPONSE 200

Пример запроса
```shell
curl --location --request PATCH 'http://localhost:8081/api/products/release' \
--header 'Content-Type: application/json' \
--data '{
    "ids": [
        "product id"
    ]
}'
```

В ОТВЕТЕ НИЧЕГО НЕ БУДЕТ
