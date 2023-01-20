# 📑 Description

It's the backend system for our [social media app project](https://github.com/incredible-four/socialmedia-app) (FE Repository)

This RESTful API was developed by using Golang and written based on Clean Architecture principles. Built with Echo as web framework, GORM as ORM, MySQL as DBMS, etc.

# About the Project

1. User can register and login
2. User can update its profile (avatar, banner, name, username, email)
3. User can deactivate its own account
4. User can CRUD contents
5. User can CRUD comments on any content
6. and many more

# ⚡Features
- CRUD (Users, Contents, Comments)
- Hashing password
- Authentication & Authorization
- Database Migration
- Automated deployment with GitHub Actions, DockerHub & AWS EC2

# 📂 Folder Structure Pattern

```
├── .github
│   └── workflows
│       └── main.yml
├── config
│   └── cloudinary.go
│   └── config.go
│   └── db.go
└── dtos
│   └── media_dto.go
└── features
│   └── comment
│   │   └── data
│   │   │   └── model.go
│   │   │   └── query.go
│   │   └── handler
│   │   │   └── handler.go
│   │   │   └── request.go
│   │   │   └── response.go
│   │   └── services
│   │   │   └── service_test.go
│   │   │   └── service.go
│   │   └── entity.go
│   └── content
│   │   └── data
│   │   │   └── model.go
│   │   │   └── query.go
│   │   └── handler
│   │   │   └── handler.go
│   │   │   └── request.go
│   │   │   └── response.go
│   │   └── services
│   │   │   └── service_test.go
│   │   │   └── service.go
│   │   └── entity.go
│   └── user
│   │   └── data
│   │   │   └── model.go
│   │   │   └── query.go
│   │   └── handler
│   │   │   └── handler.go
│   │   │   └── request.go
│   │   │   └── response.go
│   │   └── services
│   │   │   └── service_test.go
│   │   │   └── service.go
│   │   └── entity.go
└── helper
│   └── cloudinary_helper.go
│   └── jwt.go
│   └── pwd.go
│   └── response.go
└── mocks
├── .gitignore
├── cloud.env.example
├── dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── local.env.example
├── main.go
├── openapi.yaml
└── README.md
```

# 🔥 Open API

Simply [click here](https://app.swaggerhub.com/apis-docs/ALIFMUHAMADHAFIDZ23/SocialMedia-Group4/1.0.0) to see the details of endpoints we have agreed with our FE team.

<details>
  <summary>👶 Users</summary>
  
| Method      | Endpoint            | Params      | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- | ----------- | --------------------------------------- |
| POST        | /register           | -           | NO          | Register a new user                     |
| POST        | /login              | -           | NO          | Login to the system                     |
| GET         | /users/{username}   | username    | NO          | Show profile (get user & its contents)  |
| GET         | /users              | -           | YES         | Get data user (for edit profile form)   |
| PUT         | /users              | -           | YES         | Update user profile                     |
| DELETE      | /users              | -           | YES         | Deactivate user account                 |
  
</details>

<details>
  <summary>🔖 Contents</summary>
  
| Method      | Endpoint                | Params      | JWT Token   | Function                                |
| ----------- | ----------------------- | ----------- | ----------- | --------------------------------------- |
| GET         | /contents               | -           | NO          | Get all contents                        |
| GET         | /contents/{id_content}  | id_content  | NO          | Get a content by its ID                 |
| POST        | /contents               | -           | YES         | Post a new content                      |
| PUT         | /contents/{id_content}  | id_content  | YES         | Update a content                        |
| DELETE      | /contents/{id_content}  | id_content  | YES         | Delete a content                        |
  
</details>

<details>
  <summary>🗨️ Comments</summary>
  
| Method      | Endpoint                | Params      | JWT Token   | Function                                |
| ----------- | ----------------------- | ----------- | ----------- | --------------------------------------- |
| POST        | /comments/{id_content}  | id_content  | YES         | Post a new comment                      |
| DELETE      | /comments/{id_comment}  | id_comment  | YES         | Delete a comment                        |
  
</details>



# 🔗 ERD

![run](./entity-relationship-diagram.png)
# 💻 Tools & Technologies
- Backend Stacks :
  - [Golang](https://go.dev/) : Programming Language
  - [Viper](https://github.com/spf13/viper) : Environment Reader
  - [Echo](https://echo.labstack.com/) : Web Framework
  - [JWT](https://jwt.io/) : Authentication & Authorization
  - [GORM](https://gorm.io/) : ORM Library
  - [MySQL](https://gorm.io/) : Database Management System
- Documentation :
  - [Postman](https://www.postman.com/) : API Testing & Documentation
  - [Swagger](https://swagger.io/) : Open API Documentation
- Deployment & Storage :
  - [Ubuntu](https://ubuntu.com/) : Development & Deployment OS
  - [Docker](https://docker.com/) : Containerization
  - [Amazon EC2](https://aws.amazon.com/) : Cloud server
  - [Cloudinary](https://cloudinary.com/) : Store and retrieve images
  - [Cloudflare](https://www.cloudflare.com/) : SSL & Proxy

# 🛠️ How to Run Locally

- Clone it

```
$ git clone https://github.com/incredible-four/socialmedia-api/
```

- Go to directory

```
$ cd socialmedia-api
```

- Create a new database

- Rename `local.env.example` to `local.env`
- Rename `cloud.env.example` to `cloud.env`
- Adjust `local.env` & `cloud.env` as your environment settings

- Run the project

```
$ go run .
```
- Voila! 🪄

# 🤖 Author

- Alif Muhamad Hafidz :

    <a target="_blank" href="https://github.com/AlifMuhamadHafidz"><img style="vertical-align: middle;" alt="Alif" src="https://raw.githubusercontent.com/hebobibun/hebobibun/main/assets/Alif.png" width="77"></a>&ensp;

- Muhammad Habibullah :

    <a target="_blank" href="https://github.com/hebobibun"><img style="vertical-align: middle;" alt="Habib" src="https://raw.githubusercontent.com/hebobibun/hebobibun/main/assets/Habib.png" width="94"></a>&ensp;



