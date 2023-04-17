# Go Kit Sample

The "go-kit-sample" repository is a demonstration of how to utilize the Go programming language and the go-kit framework to build microservices. This repository contains code that showcases the implementation of several common microservice functionalities such as service discovery, load balancing, and distributed tracing.

The code in this repository is well-organized and structured, making it easy to understand and modify for different use cases. It is also accompanied by detailed documentation that explains each component's purpose and how to use it.

Additionally, this repository includes sample code for creating and deploying microservices using Docker and Kubernetes. This is especially useful for those interested in containerization and orchestration of microservices.

Overall, this repository serves as a useful resource for developers looking to learn about microservices, the go-kit framework, and best practices for implementing them.

## Build
```shell
docker compose up --detach
```


## APIs

### POST

#### Retail
```shell
curl --location 'http://localhost:8080/retail' \
--header 'Content-Type: application/json' \
--data '{
    "code": "aaa111",
    "qty":12

}'
```

#### Wholesale
```shell
curl --location 'http://localhost:8080/wholesale' \
--header 'Content-Type: application/json' \
--data '{
    "partner":"superstore",
    "code": "aaa111",
    "qty":12
}'
```

### GET

#### Metrics
```shell
curl --location 'http://localhost:8080/metrics'
```

## Tests

```shell
$ go test ./... 
```
