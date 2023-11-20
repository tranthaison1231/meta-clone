clean:
	rm -rfv ./bin
dev: 
	env PORT=8080 go run router/main.go
build: clean
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/api router/main.go
	zip bin/api.zip bin/api
deploy: build
	sls deploy
	
