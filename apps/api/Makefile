clean:
	rm -rfv ./bin
swagger:
	swag init
dev: swagger
	env PORT=8080 go run ./main.go
migration: 
	go run migrations/main.go
build-api:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/api ./main.go
	zip bin/api.zip bin/api
build-ws:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/ws ws/main.go
	zip bin/ws.zip bin/ws
deploy: clean build-api build-ws
	sls deploy
