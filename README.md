# Go Kit Sample


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
