BINARY=service
build:
	go build ${LDFLAGS} -o ${BINARY} cmd/web/*.go

web: build
	./${BINARY} -E dev