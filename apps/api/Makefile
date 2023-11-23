clean:
	rm -rfv ./bin
dev: 
	env PORT=8080 go run router/main.go
migration: 
	go run migrations/main.go
build-api:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/api router/main.go
	zip bin/api.zip bin/api
build-ws:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/ws ws/main.go
	zip bin/ws.zip bin/ws
deploy: clean build-api build-ws
	sls deploy
