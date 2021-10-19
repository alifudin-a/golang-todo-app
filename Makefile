tidy:
	go mod tidy

run:
	go run cmd/server/main.go

build:
	cd cmd/server; go build -o ../../bin/golang-todo-app

exec:
	./bin/golang-todo-app

mkbin:
	mkdir bin

startapp: build exec