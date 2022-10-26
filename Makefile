export COMMIT=$(shell git rev-parse HEAD)
export VERSION=development
export DATE=$(shell date)

flatten:
	swagger version
	swagger flatten spec/launcher.swagger.yaml > spec/flatten.swagger.yaml && swagger validate spec/flatten.swagger.yaml

gen: flatten
	rm -rf server/restapi || true
	swagger generate server -f spec/flatten.swagger.yaml -t server --exclude-main --principal=models.Auth

tests:
	go test -cover ./...

lint:
	gofmt -s -w ./ && golangci-lint run