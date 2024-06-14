# LRU Cache Project

This repository contains a project implementing an LRU (Least Recently Used) cache with separate components for a Golang API and a React frontend.

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

### Prerequisites

Make sure you have the following installed on your machine:

- Golang
- Node.js and npm (Node Package Manager)

### Cloning the Repository

1. Clone the repository to your local machine.
2. Navigate into the `lru-cache` directory.

### Running the Golang API

1. **Navigate to the `lru-cache-api` folder:**

    ```plaintext
    cd lru-cache-api
    ```

2. **Install dependencies:**

    ```plaintext
    go mod download
    ```

3. **Start the Golang server:**

    ```plaintext
    go run main.go
    ```

   The API server should now be running locally.

### Running the React App

1. **Navigate to the `lru-cache-app` folder:**

    ```plaintext
    cd lru-cache-app
    ```

2. **Install dependencies:**

    ```plaintext
    npm install
    ```

3. **Start the React application:**

    ```plaintext
    npm start
    ```

   The React app should now be accessible at http://localhost:3000 in your web browser.

## Contributing

Feel free to contribute to this project by forking the repository and submitting a pull request. If you have any questions or suggestions, please open an issue.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
