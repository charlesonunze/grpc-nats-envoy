# grpc-nats-envoy

json to rpc example with envoy, go, grpc, redis

This repo is a mirror of [https://github.com/charlesonunze/grpc-redis-envoy-example](https://github.com/charlesonunze/grpc-redis-envoy-example)

It replaces [redis](https://redis.io/) with [nats](https://nats.io/)

## Run

Make sure you have docker installed locally

Run the services

```bash
  docker-compose up --build
```

## Testing

POST http://localhost:1337/user/login

body {"name": "John" }

.

POST http://localhost:1337/user/balance

body {"token": "token gotten from the login" }

.

POST http://localhost:1337/transactions/up

body {"amount": 200, "token": "token gotten from the login" }

.

POST http://localhost:1337/transactions/down

body {"amount": 700, "token": "token gotten from the login" }
