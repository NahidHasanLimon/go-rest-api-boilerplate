# REST API Boilerplate in Go

This repository is a REST API project built using Go. It was initially created as a learning project for exploring REST API development in Go. Over time, the repository has been refined into a structured and reusable boilerplate for future REST API projects.

## Features
- **Standardized Project Structure:** Organized into folders such as `utils`, `routes`, `models`, `handlers`, `config` (for database and environment settings), etc.
- **Validation System:** Implemented a centralized validation message handler for simplified error messaging.
- **Scalable & Maintainable:** Designed to be modular and easy to extend for future projects.

## Folder Structure
```
├── config       # Database and environment configuration
├── handlers     # Request handlers
├── routes       # API route definitions
|── models       # Data modeling use struct
├── utils        # Utility functions
├── main.go      # Application entry point
```

## Getting Started
### Prerequisites
Ensure you have Go installed on your system.

### Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/NahidHasanLimon/go-rest-api-boilerplate
   cd go-rest-api-boilerplate
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the application:
   ```sh
   go run main.go
   ```

## Contributions
Feel free to contribute by submitting issues or pull requests!

## License
This project is licensed under the MIT License.

