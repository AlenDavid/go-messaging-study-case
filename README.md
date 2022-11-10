# A golang case-study for messaging applications

Messaging is fun!!111

## Quick start

One must have golang installed to run this applications.
Docker helps to create the developer environment by running `docker compose up` but the only external dependency is a RabbitMQ instance.

### Bootstrapping services

To bootstrap RabbitMQ:

```sh
docker compose up -d
```

To run the producer service:

```sh
go run cmd/producer/producer.go
```

To run the consumer service:

```sh
go run cmd/consumer/consumer.go
```
