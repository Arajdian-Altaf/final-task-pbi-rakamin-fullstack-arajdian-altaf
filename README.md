# Setup

Buatlah sebuah file `.env` dalam root project dengan konten seperti berikut:

```.env
DB_HOST=""
DB_USER=""
DB_PASS=""
DB_NAME=""
TOKEN_SECRET=""
```

Isi semua field `.env` sesuai dengan environment dan kebutuhan masing-masing.

# API Documentation

## Create User

- Endpoint: /users/register
- Method: POST
- Body:
```json
{
    "username": "",
    "email": "",
    "password": ""
}
```

## Login

- Endpoint: /users/login
- Method: GET
- Body:
```json
{
    "email": "",
    "password": ""
}
```

## Update user

- Endpoint: /users/:userId
- Method: PUT
- Header: `Authorization: Bearer (JWT Token)`

- Body:
```json
{
    "username": "",
    "email": "",
    "password": ""
}
```

## Delete user

- Endpoint: /users/:userId
- Method: DELETE
- Header: `Authorization: Bearer (JWT Token)`

---

## Get all photo

- Endpoint: /photos
- Method: GET
- Header: `Authorization: Bearer (JWT Token)`

## Upload photo

- Endpoint: /photos
- Method: POST
- Header: `Authorization: Bearer (JWT Token)`
- Body: 
```json
{
    "title": "",
    "caption": "",
    "photo_url": ""
}
```

## Update photo

- Endpoint: /photos/:photoId
- Method: PUT
- Header: `Authorization: Bearer (JWT Token)`

## Delete photo

- Endpoint: /photos/:photoId
- Method: DELETE
- Header: `Authorization: Bearer (JWT Token)`