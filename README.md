# My Admin Service API

A simple project starter Golang with Gin Web Framework + GORM

## Features

-   Authentication with JWT
-   CRUD Users
-   CRUD Menus
-   CRUD Roles
-   CRUD Permissions

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file, look the `.env.example` file

## API Reference

#### [ Auth ] - Register

```http
  POST /api/v1/auth/register
```

| Body       | Type     | Description   |
| :--------- | :------- | :------------ |
| `name`     | `string` | **Required**. |
| `username` | `string` | **Required**. |
| `email`    | `string` | **Required**. |
| `password` | `string` | **Required**. |

#### [ Auth ] - Login

```http
  POST /api/v1/auth/login
```

| Body       | Type     | Description   |
| :--------- | :------- | :------------ |
| `username` | `string` | **Required**. |
| `password` | `string` | **Required**. |

#

the others can be seen in Postman or Swagger

## Authors

-   [@andikaps](https://www.github.com/andikaps)
