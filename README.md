### Prerequisites

- docker, docker-compose
- go 1.20
- mysql(client)

### Setup
To generate protobuf code:

```protoc -I proto/ --go_out . --go-grpc_out . --grpc-gateway_out . proto/togo.proto```

To build:

```docker-compose build```

To start:
```docker-compose up -d```

To setup the database: 

```mysql -h localhost -P3001 --protocol=tcp -utogo --password=togo togo  < togo.sql```

### Sample API 

Via HTTP/JSON API

```curl -v http://localhost:8080/users/1/tasks -X "POST" -d "{\"name\":\"tests\"}"```

Via gRPC/Protobuf

```./bin/proto_client -addr "localhost:8081" -id 1 -name "tests"```

### Running tests

```go test ./...```


### Asynchronous Events

This project is capable of integrating with AWS SQS to handle asynchronous requests.

``` ./bin/proto_event_pusher -id <unique id> -name <task name> -queue <aws fifo endpoint> ```

``` ./bin/event_handler -queue <aws fifo endpoint> ```

### Notes

For quick setup, docker and docker-compose is used to handle dependencies.

For this project, I setup the structure to follow MVC principles. The structure is also loosely based on https://github.com/katzien/go-structure-examples/tree/master/domain-hex. 

The idea is to be able to replace the API layer(implemented JSON here, but can be swapped to SOAP, gRPC, etc), and database layer(implemented MySQL and mock test here) any time, without affecting the core logic.

For tests I used golang's native test framework for both unit tests and integration tests. The integration test verifies the API result and the MySQL database entries. 

Since the requirement is to only have 1 API endpoint, the user resource is created automatically if the user doesn't exist yet. Otherwise I would have created a separate controller/logic for the user resource. A few values also hard coded to skip config management. Some errors aren't checked because they are assumed to not fail, this isn't production grade and otherwise would've been checked for completeness.
