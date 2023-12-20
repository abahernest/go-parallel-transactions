# Golang Parallel Transactions App

## Project Description
A simple program that implements parallel execution of user transfer/transactions using chennels.



## Endpoints

[https://documenter.getpostman.com/view/11044390/2s9Ykq71Rx](https://documenter.getpostman.com/view/11044390/2s9Ykq71Rx)


## Running the app

```bash
$ go run cmd/main.go

```

## Limitations

- In-Memory data storage. (No real persistence)
- Poor Queue Implementation. The queue implemantation uses golang channels which uses an in-memory queue. This will lead to scenarios where data added to verification queue can be lost when the server crashes or restarts. A more resislient system like RabbitMQ is to be used instead.