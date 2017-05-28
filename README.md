# gRPC

![example](http://i.imgur.com/bCyxjLv.png)

A tiny test using gRPC between Golang and Ruby languages

  - [x] Server: Golang, Client: Golang
  - [x] Server: Ruby, Client: Golang
  - [x] Server: Golang, Client: Ruby
  - [x] Server: Ruby, Client: Ruby

## Using

> Both server are using "localhost:8080" by default address

Up and running one server:

```sh
# go server
go run server/server-go.go

# ruby server
ruby server/server-ruby.rb
```

then you can use any client to do some requests:

```sh
# go client:
go run client/client-go.go -id <id>

# ruby client
ruby client/client-ruby.rb <id>
```

> Both clients are using "localhost:8080" by default server address

### References

  - [gRPC QuickStart](http://www.grpc.io/docs/quickstart/)
  - [gRPC examples](https://github.com/grpc/grpc/tree/master/examples)
  - [Francesc Campoy (francesc/JustForFunc)](https://www.youtube.com/watch?v=XaMr--wAuSI)
