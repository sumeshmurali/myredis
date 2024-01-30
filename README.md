# MyRedis

This is a fun side project which is a minimal clone of redislike key store

## Features
- Support string storage
- Atomic string operations like SET, INCR, DECR is supported

## Getting started

### Prerequistes
- Go 1.20

### Starting server
Run the following command from project folder
```sh
sh start_server.sh
```

This will start the server on localhost:4000

### Accessing the server
Note: Currently uses HTTP for communication

Send a request with the following format to the web server:

path: /

method: POST

Body:
```json
{
    "command": "SET",
    "key": "mykey",
    "args": ["1"]
}
```
Note: Make sure to set the 'content-type' header to 'application/json'

## Features in plan
- Adding more data structures

