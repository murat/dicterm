build: clean
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ./bin/dicterm cmd/dicterm/main.go

clean:
	go clean
	-rm ./bin/dicterm

test:
	go test ./... -coverprofile=cover.out
	curl -Ls https://coverage.codacy.com/get.sh -o codacy.sh && \
	bash ./codacy.sh report -s --force-coverage-parser go -r cover.out -t ${CODACY_PROJECT_TOKEN}

lint:
	golangci-lint run ./... -c ./.golangci.yml