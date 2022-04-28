build: clean
	GOARCH=amd64 GOOS=linux go build -o ./bin/app cmd/app/main.go

clean:
	go clean
	-rm ./bin/app

test:
	go test ./...

coverage:
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

lint:
	golangci-lint run ./... -c ./.golangci.yml