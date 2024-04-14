# Golang CRUD API for Car Management

This is a simple CRUD (Create, Read, Update, Delete) API built with Go (Golang) for managing cars. It allows users to perform basic operations on a list of cars, including adding, retrieving, updating, and deleting car records.

## My Learning Journey with Go

I'm currently learning Go and this project is part of my journey to master the language. Here are some key aspects of my learning experience:

- **Motivation**: I was inspired to learn Go because of its simplicity, performance, and growing popularity in the industry.
- **Challenges**: Learning a new language always comes with challenges, but I've found Go's clear syntax and extensive standard library to be incredibly helpful.
- **Progress**: With each project, I'm gaining a deeper understanding of Go's features, including concurrency, error handling, and package management.
- **Community**: The Go community is amazing! I've found support and resources through online forums, documentation, and open-source projects.

## Features

- **GET /cars**: Retrieve a list of all cars.
- **POST /cars**: Add a new car to the list.
- **PUT /cars/{Company}**: Update an existing car by company name.
- **DELETE /cars/{Company}**: Delete a car by company name.

## Technologies Used

- Go (Golang)
- Gorilla Mux (for HTTP routing)
- JSON encoding/decoding

## Installation

1. Clone this repository: `git clone https://github.com/yourusername/golang-car-api.git`
2. Navigate to the project directory: `cd golang-car-api`
3. Build the project: `go build`
4. Run the executable: `./golang-car-api`

## Usage

- Make requests to the API endpoints using tools like cURL, Postman, or your preferred HTTP client.

### Example:

- To retrieve all cars:
  ```
  curl http://localhost:8080/cars
  ```

- To add a new car:
  ```
  curl -X POST -H "Content-Type: application/json" -d '{"company":"Tesla","model":"Model S","color":"Red","price":80000,"manufacturingYear":2022,"owner":{"name":"Elon Musk","age":50}}' http://localhost:8080/cars
