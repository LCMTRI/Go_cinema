# Go Cinema CRUD API

A simple CRUD REST API in Golang, re-designed to implement gRPC, with connection to MongoDB.

## System architecture:

<!-- <p align="center">
    <img alt="diagram" src="https://github.com/LCMTRI/Go_cinema/blob/main/flow_diagram.drawio" width="70%">
</p> -->

<p align="center">
    <img alt="diagram" src="https://github.com/LCMTRI/Go_cinema/blob/main/microservice.jpg" width="50%">
</p>

<p align="center">
    <img alt="diagram" src="https://github.com/LCMTRI/Go_cinema/blob/main/go_cinema.jpg" width="70%">
</p>

## Database connection

Before running the program, consider setting up MongoDB, then changing the config (URI, databases' names) in <code>database/connection.go</code>

## Run

First we run 2 __gPRC Client Servers__:

```
> go run movie/movie.go
> go run user/user.go
```

Then we run our main __Gateway Server__ which will interact with public REST API:

```
> go run gateway/main.go
```