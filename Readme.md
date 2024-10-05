## STRUCTURE API

* success response
```json
{
  "code": 200,
  "message": "success",
  "data": ""
}
```

* error response
```json
{
  "code": 500,
  "message": "failed to store",
  "data": ""
}
```
---

### Product

* Get Products

```
/api/products
```

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "3f5b407e-eb3d-4742-b0f1-56dc61ce5fca",
    "name": "B",
    "sku": "02",
    "supplier": {
      "id": "67fb6dcb-4dee-4fe9-9867-59d04ec0b84b",
      "name": "PT. Hamukti Ramono Rizki"
    },
    "brand": {
      "id": "93ee356b-5272-48ea-8c45-67c99a35a3a9",
      "name": "Adidas"
    }
  }
}
```

* Get Product

```
/api/products/:id
```

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": "aec21e50-32ab-4e89-9a6e-f5716f58627e",
      "name": "A",
      "sku": "01",
      "supplier": {
        "id": "1d78b253-6901-4959-99c5-015a01a6dc9b",
        "name": "PT. Trijoyo Putro Mardiantoro"
      },
      "brand": {
        "id": "a15a1ae6-5fc8-46b0-a692-0df0d860ec7c",
        "name": "Nike"
      }
    },
    {
      "id": "3f5b407e-eb3d-4742-b0f1-56dc61ce5fca",
      "name": "B",
      "sku": "02",
      "supplier": {
        "id": "67fb6dcb-4dee-4fe9-9867-59d04ec0b84b",
        "name": "PT. Hamukti Ramono Rizki"
      },
      "brand": {
        "id": "93ee356b-5272-48ea-8c45-67c99a35a3a9",
        "name": "Adidas"
      }
    }
  ]
}
```