all: build

DB_USER=postgres
APP=crud-app
APP_EXECUTABLE="./out/$(APP)"
DB_NAME=$(APP)
TEST_DB_NAME="crud_app_test"

build-deps:
	glide install

compile:
	go build -o $(APP_EXECUTABLE)

build: build-deps compile

db.setup: db.create 

db.create:
	createdb -O$(DB_USER) -Eutf8 $(DB_NAME)

db.drop:
	dropdb --if-exists -U$(DB_USER) $(DB_NAME)

db.reset: db.drop db.create 

testdb.setup: testdb.drop testdb.create 

testdb.create: testdb.drop
	createdb -O$(DB_USER) -Eutf8 $(TEST_DB_NAME)

testdb.drop:
	dropdb --if-exists -U$(DB_USER) $(TEST_DB_NAME)



