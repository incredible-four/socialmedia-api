# Description

It's a backend system for our [social media app project](https://github.com/incredible-four/socialmedia-app). User can post a content and comment on other's posts. This REST API was developed using Golang and written based on Clean Architecture principles with Echo Framework as web framework, GORM as ORM and MySQL as DBMS.

# ⚡Features
- CRUD (Users, Contents, Comments)
- Hashing password
- Authentication & Authorization
- Database Migration

# 🔥 Open API

Here's the list endpoints we have agreed on :

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

![run](./ERD.png)
# 💻 Built with

- [Golang](https://go.dev/) : Programming Language
- [Viper](https://github.com/spf13/viper) : Environment Reader
- [Echo](https://echo.labstack.com/) : Web Framework
- [JWT](https://jwt.io/) : Authentication & Authorization
- [GORM](https://gorm.io/) : ORM Library
- [MySQL](https://gorm.io/) : Database Management System
- [AWS](https://aws.amazon.com/) : Cloud Server

# 🛠️ Installation Steps

[ Available soon ]

# 👑 Author

- [Alif Muhamad Hafidz](https://github.com/AlifMuhamadHafidz)
- [Muhammad Habibullah](https://github.com/hebobibun)



