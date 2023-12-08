# Run the application locally
run:
	go run main.go

set-pre-commit:
	pre-commit install --hook-type pre-commit --config .pre-commit-config.yaml; \
	pre-commit install --hook-type pre-push --config .pre-push-config.yaml

coverage:
	go test -v -coverprofile cover.out ./...
	go tool cover -html cover.out -o cover.html

build:
	env GOOS=linux GOARCH=amd64 go build -o bin/cdi-generate-token-onetime-v1
