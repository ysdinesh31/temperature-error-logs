# Temperature Error Logs API

## Overview

This project is a Dockerized Golang API designed to handle temperature readings and store improperly formatted data in MongoDB. The API includes the following endpoints:

- **POST /temp**: Accepts temperature data in a specified format. Returns whether the temperature is above a threshold and formats the timestamp if needed. If the data format is incorrect, it returns an error.
- **GET /errors**: Retrieves all incorrectly formatted data strings.
- **DELETE /errors**: Clears the error logs.

The application is deployed on an AWS EC2 instance and uses MongoDB for persistent storage.

## Running Locally

### Prerequisites

- **Docker**: Ensure Docker is installed on your machine.
- **Docker Compose**: Ensure Docker Compose is installed.

### Steps to Run

1. **Clone the Repository**

   ```bash
   git clone https://github.com/ysdinesh31/temperature-error-logs
2. **Navigate to the Project Directory**

   ```bash
    cd temperature-error-logs
3. **Build and Run with Docker Compose**

   ```bash
    docker-compose up --build
