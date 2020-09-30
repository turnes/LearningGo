# Goroutines

1. goroutine-basic
1. racecondition
1. mutex
1. atomic


## Check race condition

`go run -race main.go`

> Found X data race(s)

## Race condition

|project  |race condition  |
|---------|----------------|
|basic    |       No       |
|race condition     |       Yes         |
|mutex     |      No          |
|atomic     |      No          |
