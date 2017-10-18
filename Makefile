all: build testdb.create testdb.migrate test

DB_USER=postgres
APP=crud-app
APP_EXECUTABLE="./out/$(APP)"
DB_NAME=$(APP)
TEST_DB_NAME="crud_app_test"
UNIT_TEST_PACKAGES=$(shell glide novendor)

build-deps:
	glide install

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

build: build-deps compile

install:
	go install ./...

db.setup: db.create 

db.create:
	createdb -O$(DB_USER) -Eutf8 $(DB_NAME)

db.drop:
	dropdb --if-exists -U$(DB_USER) $(DB_NAME)

db.migrate:
	ENVIRONMENT=development $(APP_EXECUTABLE) migrate

db.reset: db.drop db.create 

test:
	ENVIRONMENT=test go test $(UNIT_TEST_PACKAGES) -p=1

testdb.setup: testdb.drop testdb.create 

testdb.create: testdb.drop
	createdb -O$(DB_USER) -Eutf8 $(TEST_DB_NAME)

testdb.migrate:
	ENVIRONMENT=test $(APP_EXECUTABLE) migrate

testdb.drop:
	dropdb --if-exists -U$(DB_USER) $(TEST_DB_NAME)



