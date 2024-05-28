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

## Ethereum Web Token (EWT) Authentication

This project implements Ethereum Web Token (EWT) based authentication for HTTP handlers. The implementation includes creating, validating, and managing EWTs.

### Overview

The project contains functionality to:
- Sign messages with an Ethereum private key.
- Verify signatures.
- Create tokens that include a signature response and an expiry date.
- Validate tokens and use them for HTTP authentication.

## Package `ewt`

### Types

#### `SignatureResponse`

Represents the structure of the signature response.

```go
type SignatureResponse struct {    
	Address string `json:"address,omitempty"`    
	Msg     string `json:"msg,omitempty"`    
	Sig     string `json:"sig,omitempty"`    
	Version string `json:"version,omitempty"`
}



## Installation

To run the project, follow these steps:

1. Create a MySQL database named `projectmanager`:

   ```bash
   mysql -u root -p
   Enter password: [your_password]
   mysql> CREATE DATABASE projectmanager;

2. Run on terminal `make run`
