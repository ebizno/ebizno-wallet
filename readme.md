# E-Bizno

**E-Bizno** _is a web application Buying and Selling Houses and Land is a system that
will allow the purchase and sale of real estate in Angola, a system developed for the web with the
objective of facilitating the adhesion of properties in Angola._

**E-Bizno** é uma aplicação web Compra e Venda de Casas e Terrenos é um sistema que
permitirá a compra e venda de imóveis em Angola, sistema desenvolvido para a web com a
objectivo de facilitar a adesão de imóveis em Angola

**E-Bizno** _is a web application that will allow the purchase and sale of real estate in Angola,
a system developed for the web with the objective of facilitating the adhesion of real estate in Angola._

**E-Bizno** uma aplicação web que vai permitir a compra e venda de imóveis em Angola,
um sistema desenvolvido para a web com o objetivo de juntar imóveis em Angola.


```shell
echo "module.exports = {extends: ['@commitlint/config-conventional']}" > commitlint.config.js

yarn husky add .husky/commit-msg 'yarn commitlint --edit $1'

#or

npm husky add .husky/commit-msg 'npm commitlint --edit $1'
```

```shell
export PATH=$(go env GOPATH)/bin:$PATH
swag init 
or
swag init -g adapter/user/rest/routes/user_router_api.go --output adapter/user/rest/routes/docs
```

ACESS PORT 👇
### User
- rest=2003
- grpc=3003

### Wellet
- rest=2004
- grpc=3004

## Address
- rest=2005
- grpc=3005

### Cart
- rest=2001
- grpc=3001

### Client
- rest=2002
- grpc=3002

### Product
- rest=2000
- grpc=3000


<h3 align="center"> 
    🚀 Documentation the API´s 🚀
</h3>

## Authentication

```http request
POST 127.0.0.1:2003/user/auth
Content-Type: application/json
{
    "email": "pl000000@gmail.com",
    "password": "1234"
}
or
{
    "bi": "000000000LA000",
    "password": "1234"
}
```
### response
```json
{
    "data": {
        "IdUser": "d7ea7136-a96a-4278-af2e-c7603e4e661d",
        "FirstName": "first",
        "LastName": "last",
        "telephone": "900000000",
        "bi": "000000000LA000",
        "email": "pl0000000@gmail.com",
        "password": "$2a$10$xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        "state": true
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "wellet": 0
}
```

## Authorization

| Parameter | Type | Description                  |
| :--- | :--- |:-----------------------------|
| `Authorization` | `Bearer` | **Required**. Your API Token |

## Status Codes

returns the following status codes in its API:

| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 201 | `CREATED` |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND` |
| 500 | `INTERNAL SERVER ERROR` |


## Clean Architecture

## How run application
We use Docker so that all the services we use are available.

- Make the project clone
- Having docker installed on your machine just run:
`docker-compose up -d`


To continue ...