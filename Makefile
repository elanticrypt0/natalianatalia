BINARY_NAME=app
BUILD_DIR=./build
HTMLX_DIR=./public/assets/js

go-install-deps:
	go install github.com/cosmtrek/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go mod tidy

build:
	# create directories
	mkdir -p ${BUILD_DIR}
	mkdir ${BUILD_DIR}/_db
	mkdir ${BUILD_DIR}/logs
	mkdir ${BUILD_DIR}/public

	# copy the config files and seeds
	cp -r ./config ./build/config
	cp -r ./seeds ./build/seeds
	
	# templates generator
	templ generate

	# compile into binary file
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} -ldflags "-w -s"
	chmod +x ${BINARY_NAME}
	mv ${BINARY_NAME} ${BUILD_DIR}

templates:
	templ generate

dev:
	templ generate
	air

run:
	templ generate
	go run .

test:
	go test ./tests

clean:
	go clean
	rm -rf ${BUILD_DIR}

htmlx-install:
	wget https://unpkg.com/htmx.org/dist/htmx.min.js -P ${HTMLX_DIR}

htmlx-update:
	rm -rf ${HTMLX_DIR}/*
	wget https://unpkg.com/htmx.org/dist/htmx.min.js -P ${HTMLX_DIR}