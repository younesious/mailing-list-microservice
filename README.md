# Mailing List Microservice

## Description

This project is a microservice-based mailing list management system built with Go, gRPC, and Docker. It includes an HTTP API for handling JSON requests and a gRPC API for more efficient communication between microservices. The system uses SQLite as its database.

## Project Structure

```sh
├── db
│   └── db.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── grpc_api
│   ├── grpc_client
│   │   ├── Dockerfile
│   │   └── grpc_client.go
│   └── grpc_server
│       └── grpc_server.go
├── json_api
│   └── handlers.go
├── LICENSE
├── mailinglist.db
├── proto
│   ├── email_grpc.pb.go
│   ├── email.pb.go
│   └── email.proto
├── README.md
└── server
    ├── Dockerfile
    └── server.go

7 directories, 16 files

```

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Running the Project

1. Clone the repository:

   ```sh
   git clone https://github.com/younesious/mailing-list-microservice.git
   cd mailing-list-microservice
   ```

2.  Build and start the services using Docker Compose:

    ```sh
    docker compose up --build
    ```

3.  The HTTP server will be available on `http://localhost:8080` and the gRPC server on `http://localhost:8081`.

### Interacting with the Services

#### HTTP API Endpoints

-   Add Email: `POST /email/add`
-   Get Email: `GET /email/get`
-   Update Email: `PUT /email/update`
-   Delete Email: `DELETE /email/delete`
-   Get Email Batch: `GET /email/batch`

#### gRPC API

The gRPC API can be accessed using any gRPC client. Refer to the `proto/email.proto` file for the service definitions.

### Building the Project

To build the project manually:

+  Build the server:

    ```sh
    docker build -t myserver -f server/Dockerfile .
    ```
+  Build the client:

    ```sh
    docker build -t myclient -f grpc_api/grpc_client/Dockerfile .
    ```

### Running the Project

To run the gRPC server and client:

```sh
docker compose up --build
```

License
-------

This project is licensed under the MIT License - see the [LICENSE](https://github.com/younesious/mailing-list-microservice/blob/master/LICENSE) file for details.

### Contributing

Feel free to contribute and I'll be happy to see you :)
