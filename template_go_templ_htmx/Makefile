templ:
	@templ generate

build: templ
	@GOOS=linux GOARCH=amd64 go build -o bin/app .

test:
	@go test -v ./...
	
dev: templ
	@DB_NAME="myDB" go run .



