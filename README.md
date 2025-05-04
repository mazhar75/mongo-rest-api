---

```markdown
# Mongo REST API

This is a basic RESTful API built with **Go**, using **Gorilla Mux** for routing and **MongoDB** as the database. It performs standard **CRUD (Create, Read, Update, Delete)** operations on a `User` collection. The API is tested using **Postman**.

## 🛠 Tech Stack

- **Go (Golang)**
- **MongoDB Atlas (Cloud DB)**
- **Gorilla Mux** (Routing)
- **Postman** (API Testing)

## 📁 Project Structure

```

mongo-rest-api/
├── main.go
└── go.mod

````

## 📦 Features

- Connects to MongoDB Atlas
- Performs CRUD operations:
  - Create a new user
  - Get all users
  - Get a user by ID
  - Update a user
  - Delete a user

## 📄 User Model

```json
{
  "id": "ObjectID",
  "name": "string",
  "email": "string"
}
````

## 🚀 API Endpoints

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| POST   | `/users`      | Create a new user |
| GET    | `/users`      | Get all users     |
| GET    | `/users/{id}` | Get a user by ID  |
| PUT    | `/users/{id}` | Update a user     |
| DELETE | `/users/{id}` | Delete a user     |

## 🧪 Testing with Postman

1. **Start the server**
   Run this command:

   ```bash
   go run main.go
   ```

   Server runs at `http://localhost:8080`.

2. **Test Endpoints in Postman**:

   * **POST** `/users`: Send raw JSON like:

     ```json
     {
       "name": "John Doe",
       "email": "john@example.com"
     }
     ```
   * **GET** `/users`: View all users.
   * **GET** `/users/{id}`: Replace `{id}` with actual MongoDB ObjectID.
   * **PUT** `/users/{id}`: Send updated JSON.
   * **DELETE** `/users/{id}`: Delete the specified user.

## 🧰 Prerequisites

* Go installed (`go version`)
* A MongoDB Atlas cluster URI
* Git and Postman

## 🧑‍💻 Author

Developed by Mazharul.

