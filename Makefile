BINARY=service
build:
	go build ${LDFLAGS} -o ${BINARY} cmd/web/*.go

test:
	go test -race ./...

web: build
	./${BINARY} -E dev