all: build

APP=crud-app
APP_EXECUTABLE="./out/$(APP)"

build-deps:
	glide install

compile:
	go build -o $(APP_EXECUTABLE)

build: build-deps compile

