#env GOOS=linux GOARCH=arm go build -v github.com/path/to/your/app

env GOOS=linux GOARCH=arm go build -o fineractmiddleware.bin -v  cmd/main.go
