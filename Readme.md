# STRUCTURE API

## Response
#### success response
```json
{
  "code": 200,
  "message": "success",
  "data": ""
}
```

#### error response
```json
{
  "code": 500,
  "message": "failed to store",
  "data": ""
}
```
---
#### Table of response

| Code | Description      | Model               |
|------|------------------|---------------------|
| 200  | Success          | get, update, delete |
| 201  | Created          | create              |
| 401  | Unauthorized     | protected access    |
| 433  | Validation Error | -                   |
| 500  | Failed           | -                   |


---

# Endpoint API

## Authentication


```
POST /api/auth
```

#### body

| header          |      value       |
|:----------------|:----------------:|
| Content-Type    | application/json |
| Accept          | application/json |
| Authorization   |        -         |

```json
{
  "username": "xaviera@gmail.com",
  "password": "123321"
}
```
#### result
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InhhdmllcmFAZ21haWwuY29tIiwiZXhwIjoxNzI4MjAyMjI0LCJpZCI6IjgyNTY3MGI4LTg2MzMtNDg4Yy1iNzlkLTdmMjVkYTUwZjkxNSIsIm5hbWUiOiJYYXZpZXJhIFB1dHJpIn0.2JrnP0OSBfCAR2N_GRSaiv_Zvju7Gb4K5ufuNe-t5b0"
  }
}
```

## Product

| header        |         value         |
|:--------------|:---------------------:|
| Content-Type  |   application/json    |
| Accept        |   application/json    |
| Authorization |   Bearer {{token}}    |

### Get Products

```
GET /api/products
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

### Get Product

```
GET /api/products/:id
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

## Brand

| header        |      value       |
|:--------------|:----------------:|
| Content-Type  | application/json |
| Accept        | application/json |
| Authorization | Bearer {{token}} |

### Get Brand

```
GET /api/brands
```

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": "a15a1ae6-5fc8-46b0-a692-0df0d860ec7c",
      "name": "Nike",
      "description": ""
    },
    {
      "id": "7caba839-f86d-44c8-9e3f-ea87f1fd20f3",
      "name": "Samsung",
      "description": ""
    },
    {
      "id": "22343448-228c-4c92-a337-6abaecc95984",
      "name": "Huawei",
      "description": ""
    },
    {
      "id": "4476bc09-1f79-4a74-af52-ce5fc426482b",
      "name": "Xiaomi",
      "description": ""
    },
    {
      "id": "f6dbf726-a1ff-4e90-8451-996e618a9309",
      "name": "Acer",
      "description": ""
    },
    {
      "id": "a4b5e24e-9320-42e5-8483-fa01ab6fbd0e",
      "name": "Asus",
      "description": ""
    },
    {
      "id": "741955db-e1a5-4fec-acd2-10c2b2616036",
      "name": "HP",
      "description": ""
    },
    {
      "id": "381cd307-4cbe-401e-87d4-6ee886f60edd",
      "name": "Toshiba",
      "description": ""
    },
    {
      "id": "93ee356b-5272-48ea-8c45-67c99a35a3a9",
      "name": "Adidas",
      "description": ""
    }
  ]
}
```

### Get Brand

```
GET /api/brands/:id
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

### Create Brand

```
POST /api/brands
```

#### body


```json
{
  "name": "Samsung"
}
```
#### result
```json
{
  "code": 201,
  "message": "created successfully",
  "data": ""
}
``` 