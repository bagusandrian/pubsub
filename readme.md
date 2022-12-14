# Dummy Pubsub

This repo is for dummy aplication impementation of pubsub redis.

## Prerequisite
1. minimum Redis v 5.X.X
2. minimum Golang V1.15.X

## Installation

1. Clone this repo on your `gopath`
2. Pull depedency this repo using `go get` or `go mod`. Comeback to your preference. Or u can search how to get depedency on golang. 
3. Run aplication using command
```
Go run main.go
```

## How to check pubsub
If aplication running well, will showing
```
$ go run main.go 
2022/12/14 13:31:37 Everything is fine!!!
example: subscribe 1
```
thats mean application is running and already subscribe on channel. 
Another Way to check, u can check on redis-cli with this command: 
```
redis-cli 
pubsub numsub [name_channel]
```
will showing number of subscriber on that channel. 