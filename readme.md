# Multi-Purpose Boilerplate for Go

This project is a multi-purpose boilerplate designed to jumpstart your Go development. If you're tired of setting up the same basic configurations for every new project, this boilerplate is for you! It provides a solid foundation with essential features already implemented, so you can focus on building your application logic.

## Features

- **Routing:** Uses Gin for fast and efficient routing. Includes example routes for user creation (`POST /users`) and retrieval (`GET /users/:id`).
- **Dependency Injection:** Uses `wire` for clean, testable code. The `InitWiring()` function sets up dependencies for each layer, and this pattern can be applied to improve your API.
- **Configuration:** Leverages Viper for flexible configuration management.
- **Logging:** Implements structured logging with Zerolog for easy debugging and monitoring.
- **Error Handling:** Includes robust error handling to gracefully manage unexpected situations.
- **Go Version:** Built with Go 1.23.

## Getting Started

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/luthfiarsyad/multi-purpose-boiler.git
    cd multi-purpose-boiler
    ```

2.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Run The App**

        ```bash
        go run cmd/main.go

    The application will be running on port 8080.

## **Usage**

### Example Routes

- **Create User:**

  ```bash
  POST /users
  ```

  **Request Body (JSON):**

  ```JSON
  {
  "name": "John Doe",
  "email": "[email address removed]"
  // ... other user fields
  }
  ```

  **Response**

  ```JSON
  {
    "message": "User created" // Or an error message
  }

  ```

- **Get User:**

  ```bash
  GET /users/:id
  ```

  _Replace :id with the actual user ID._

  **Response**

  ```JSON
    {
        "id": 1,
        "name": "John Doe",
        "email": "[email address removed]"
        // ... other user fields
    }
  ```

  Or an error message if the user is not found.

## Configuration

Configuration is managed using Viper. You can configure the application using environment variables, configuration files (e.g., config.yaml, config.json), or command-line flags. See the Viper documentation for more details.

## Logging

Logging is done with Zerolog. Logs are formatted as JSON, making them easy to parse and analyze. The boilerplate provides a pre-configured logger instance that you can use throughout your application.

## Error Handling

The boilerplate includes error handling mechanisms to catch and manage errors gracefully. Errors are logged using Zerolog, and appropriate error responses are returned to the client.

## Dependency Injection

wire is used for dependency injection. The InitWiring() function in wire.go sets up all the necessary dependencies. This makes your code more modular, testable, and maintainable.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

License
This project is licensed under the MIT License - see the LICENSE 1 file for details.
