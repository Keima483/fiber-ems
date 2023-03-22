# Employee Management System in Go Fiber

This is an attempt to make an employee management system so that i can get familiar with the go language and the fiber framework. As you can guess the code is written in go lang with the fiber framework and gorm as the ORM, The database used is MySQL. I also tried to use JWT authentication as much as i can.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

To run this project, you will need to have the following installed:

* Go version 1.16 or higher
* MySQL
* Git

### Installing

1. Clone the repository:
```
git clone https://github.com/Keima483/fiber-ems.git
```
2. Navigate to the project directory:
```
cd fiber-ems
```
3. Install the dependencies:
```
go mod tidy
```
4. Create a `app.env` file based on `app.env.example` and fill in your MySQL database credentials.

### Running the Project

1. Run the project:
```
go run main.go
```
2. Open your web browser and navigate to http://localhost:3000/. You should see a message indicating that the server is running.

## Usage
* `POST /api/v1/company/login`: Created a JWT for the by validating email and password 
* `POST /api/v1/company/signup`: Creates a new Company and also returns a JWT
* `GET /api/v1/employee`: Get All employees of the company whose token is given
* `POST /api/v1/employee`: Adds an employee to the company whose token is given
* `PUT /api/v1/employee`: Updates the user if he/she belongs to the company whose token is given 
* `DELETE /api/v1/employee/:id`: Delete the user whose id is provided if he/she biltongs to the company whose token is given

### Request Examples

Login
```javascript
POST http://localhost:3000/api/v1/company/login
Content-Type: application/json

{
  "email": "example@example.com",
  "password": "password"
}
```

Get all employees
```javascript
GET http://localhost:3000/api/v1/employee
Authorization: Bearer <token>
```

Add new employees
```javascript
POST http://localhost:3000/api/v1/employee
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "example K",
  "age": 21,
  "department": "SDE",
  "email": "example@examle.com",
  "baseSalary": 2000000
}
```

### Built With
* [Go](https://go.dev/)
* [Fiber](https://gofiber.io/)
* [GORM](https://gorm.io/)
* [MySQL](https://www.mysql.com/)
