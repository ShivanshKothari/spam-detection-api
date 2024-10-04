# User Authentication API

A RESTful API for user authentication and management.

## Table of Contents

* [Introduction](#introduction)
* [Features](#features)
* [Requirements](#requirements)
* [Installation](#installation)
* [Usage](#usage)
* [API Endpoints](#api-endpoints)
* [Models](#models)
* [Contributing](#contributing)
* [License](#license)

## Introduction

This API provides a simple way to manage user authentication and data. It uses a MongoDB database to store user information and provides endpoints for registration, login, and user management.

## Features

* User registration with phone number and password
* User login with phone number and password
* User data management (e.g. updating user information)
* Password hashing and verification

## Requirements

* Go 1.14 or higher
* MongoDB 4.2 or higher
* Gin framework
* JWT library

## Installation

1. Clone the repository: `git clone https://github.com/your-username/user-auth-api.git`
2. Install dependencies: `go get -u ./...`
3. Create a MongoDB database and add the connection string to `config/config.go`
4. Run the API: `go run main.go`

## Usage

Use a tool like `curl` or a REST client to interact with the API.

## API Endpoints

### Registration

* `POST /api/register`: Register a new user with phone number and password

### Login

* `POST /api/login`: Login an existing user with phone number and password

### User Management

* `GET /api/user`: Get the current user's information
* `PUT /api/user`: Update the current user's information
* `DELETE /api/user`: Delete the current user's account

## Models

* `User`: Represents a user with a unique ID, phone number, and password

## Contributing

Contributions are welcome! Please submit a pull request with your changes.

## License

This project is licensed under the MIT License. See `LICENSE` for details.