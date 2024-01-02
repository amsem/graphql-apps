# GraphQL Player API and GitHub License Client

This repository consists of two separate components: a GraphQL server for managing player data and a GitHub License client that queries license information for a specific key.

## Server (GraphQL Player API)

### Getting Started

To run the GraphQL server:

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/graphql-server.git
    ```

2. Navigate to the server directory:

    ```bash
    cd graphql-server
    ```

3. Install dependencies:

    ```bash
    go get -u github.com/graphql-go/graphql
    go get -u github.com/graphql-go/handler
    ```

4. Run the server:

    ```bash
    go run main.go
    ```

   The GraphQL server will be accessible at `http://localhost:8000/graphql`.

### GraphQL Queries

The server supports the following GraphQL query:

#### Get All Players

Retrieve information about all players:

```graphql
query {
  players {
    id
    name
    highscore
    isonline
    location
    levelsunlocked
  }
}
```

### Player Object

The GraphQL schema defines a `Player` object with fields like `id`, `name`, `highscore`, `isonline`, `location`, and `levelsunlocked`.

### Technologies Used

- [Go](https://golang.org/)
- [GraphQL-Go](https://github.com/graphql-go/graphql)
- [GraphQL-Go/Handler](https://github.com/graphql-go/handler)

### Contributing

Contributions are welcome! Feel free to fork the repository and submit pull requests. Issues and feature requests are also appreciated.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Client (GitHub License Client)

### Getting Started

To run the GitHub License client:

1. Clone the repository:

    ```bash
    git clone https://github.com/amsem/graphql-apps.git
    ```

2. Navigate to the client directory:

    ```bash
    cd graphql-apps
    ```

3. Install dependencies:

    ```bash
    go get -u github.com/machinebox/graphql
    ```

4. Run the client:

    ```bash
    go run main.go
    ```

### GitHub License Query

The client queries GitHub License information for a specific key (e.g., "MIT").

### Technologies Used

- [Go](https://golang.org/)
- [MachineBox GraphQL Client](https://github.com/machinebox/graphql)

### Setting up GitHub Token

Make sure to set your GitHub token as an environment variable:

```bash
export GITHUB_TOKEN=your-github-token
```

### Contributing

Contributions are welcome! Feel free to fork the repository and submit pull requests. Issues and feature requests are also appreciated.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
