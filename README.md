# Project Manager API

Project Manager API is a RESTful API for managing tasks and users within projects.

## Overview

This project consists of various components:

- **store.go**: Contains SQL queries associated with the `Store` struct.
- **types.go**: Defines structs for database entities and JSON unmarshalling.
- **tasks.go**: Handles task table queries and endpoints (`/tasks`).
- **user.go**: Handles user table queries and endpoints (`/users`).
- **utils.go**: Contains utility functions such as password hashing and JSON writing.
- **db.go**: Responsible for connecting to the database and initializing tables.
- **api.go**: Initializes and runs the server.
- **auth.go**: Handles JWT authentication.
- **config.go**: Retrieves environment variables.

## Installation

To run the project, follow these steps:

1. Create a MySQL database named `projectmanager`:

   ```bash
   mysql -u root -p
   Enter password: [your_password]
   mysql> CREATE DATABASE projectmanager;

2. Run on terminal `make run`
