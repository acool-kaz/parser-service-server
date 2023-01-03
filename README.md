# README

As part of this repository, the web-scrapper service has been implemented

Links to another services:
- [CRUD Service](https://www.github.com/acool-kaz/post-crud-service-server)
- [API Gateway](https://www.github.com/acool-kaz/api-gateway-service)


# Complete task

This project contains three services:

- Service 1: Collects 50 pages of posts from the open API at https://gorest.co.in/public/v1/posts and stores the collected data in a PostgreSQL database. Optional: Data collection can be performed in multiple threads.

- Service 2: Implements CRUD logic for the collected posts, including the ability to retrieve multiple posts, retrieve a specific post, delete a post, modify a post, and create a post.

- Service 3: Acts as an API gateway and provides methods to perform the operations of Service 1 and Service 2. Optional: Service 3 can test coverage and interact with other services via gRPC.

# Tools

- Golang
- gRPC
- PostgreSQL
- Protobuf


# Project structure

```code
.
├── cmd
│   └── main.go
├── config.json
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── app.go
│   ├── config
│   │   └── config.go
│   ├── delivery
│   │   ├── grpc
│   │   │   └── parser
│   │   │       └── handler.go
│   │   └── http
│   │       └── handler.go
│   ├── models
│   │   ├── parser.go
│   │   └── post.go
│   ├── repository
│   │   ├── post.go
│   │   ├── postgres.go
│   │   └── repository.go
│   └── service
│       ├── post.go
│       └── service.go
├── Makefile
├── migrations
│   ├── down.sql
│   └── up.sql
├── pkg
│   └── parser
│       ├── parser_grpc.pb.go
│       └── parser.pb.go
├── proto
│   └── parser.proto
└── README.md

15 directories, 22 files
```