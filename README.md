# Spam Detection API

A RESTful API for detecting spam phone numbers.

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

This API provides a simple way to detect spam phone numbers. It uses a MongoDB database to store user data, contact information, and spam reports.

## Features

* User registration and login
* Contact management
* Spam reporting
* Search by phone number
* Spam likelihood calculation

## Requirements

* Go 1.14 or higher
* MongoDB 4.2 or higher
* Gin framework
* JWT library

## Installation

1. Clone the repository: `git clone https://github.com/your-username/spam-detection-api.git`
2. Install dependencies: `go get -u ./...`
3. Create a MongoDB database and add the connection string to [config/config.go](cci:7://file:///d:/WebDevelopmentMERN/go/spam-detection-api/config/config.go:0:0-0:0)
4. Run the API: `go run main.go`

## Usage

Use a tool like `curl` or a REST client to interact with the API.

## API Endpoints

### Authentication

* `POST /api/register`: Register a new user
* `POST /api/login`: Login an existing user

### Contact Management

* `GET /api/searchbyname`: Search for contacts by name
* `GET /api/searchbynumber`: Search for contacts by phone number

### Spam Reporting

* `POST /api/mark-spam`: Mark a phone number as spam

### Spam Likelihood

* `GET /api/searchbynumber`: Get the spam likelihood of a phone number

## Models

* [User](cci:2://file:///d:/WebDevelopmentMERN/go/spam-detection-api/models/user.go:6:0-12:1): Represents a user with a unique ID, name, phone number, and email
* [Contact](cci:2://file:///d:/WebDevelopmentMERN/go/spam-detection-api/models/contact.go:6:0-11:1): Represents a contact with a unique ID, user ID, name, and phone number
* [Spam](cci:2://file:///d:/WebDevelopmentMERN/go/spam-detection-api/models/spam.go:6:0-9:1): Represents a spam report with a unique ID and phone number

## Contributing

Contributions are welcome! Please submit a pull request with your changes.

## License

This project is licensed under the MIT License. See `LICENSE` for details.