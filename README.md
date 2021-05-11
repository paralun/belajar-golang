# Belajar Golang

#### Compile & Run
```
$ go build hello.go
$ ./hello

$ go run hello.go
``` 

#### Go Modules
```
$ go mod init nama-module
$ go mod init golang-basic
$ go mod init github.com/paralun/golang-basic

membuat tag di git
$ git tag v1.0.0
$ git push origin v1.0.0

menambah dependency
$ go get nama-module
$ go get github.com/paralun/golang-basic

upgrade dependency
$ go get
```

#### Go Testing
```
$ go test -v
$ go test -v ./...
$ go test -v -run=TestHello (function testing)
```