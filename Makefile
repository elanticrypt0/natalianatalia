BINARY_NAME=app
BUILD_DIR=./build

all: build

deps:
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
	
	# compile into binary file
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} -ldflags "-w -s"
	chmod +x ${BINARY_NAME}
	mv ${BINARY_NAME} ${BUILD_DIR}


watch:
	air

run:
	go run .

test:
	go test ./tests

clean:
	go clean
	rm -rf ${BUILD_DIR}