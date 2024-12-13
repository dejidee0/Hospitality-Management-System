# Makefile
# Makefile is used to make runnig commands easy

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Define the command as a variable
# Once you use .PHONY for server, make will always execute the go run command, even if there are no file changes.
.PHONY: server build migration migrateup migratedown migrateforce
# run default command, just by typing make
# default: all
default: server

# run specified commands
# all: server tailwind # this will not work as intended thouh, since server will keep running,

# Target to run the Tailwind command
# make tailwind
# tailwind:
# 	$(TAILWIND_CMD)

# Start the Go server
# make server
server:
	go run main.go

# build a single binary
build:
	go build -o ../bin/hms main.go && cp env.example ../bin && cp .env ../bin

migration:
	migrate create -dir database/migrations -ext sql -seq $(NAME)


# 	migrate -path database/migrations -database "sqlite3://data/hms.db?sslmode=disable" -verbose
# migrateup:
# 	migrate -path database/migrations -database "sqlite3:///$(PWD)/data/hms.db" -verbose up

# migratedown:
# 	migrate -path database/migrations -database "sqlite3:///$(PWD)/data/hms.db" -verbose down


# DATABASE_HOST=SQL1001.site4now.net
# # default is 1433 for localhost
# DATABASE_PORT=
# DATABASE_NAME=db_ab0568_findpeace
# DATABASE_USERNAME=db_ab0568_findpeace_admin
# DATABASE_PASSWORD=Findpeace1234

# load_env:
# 	bash -c 'source .env && echo $DATABASE_NAME'
migrateup:
	migrate -path database/migrations -database "sqlserver://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}" -verbose up
# source .env && migrate -path database/migrations -database "mssql://<username>:<password>@<host>:<port>/<database>?param1=value1&param2=value2" -verbose up

migratedown:
	migrate -path database/migrations -database "sqlserver://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}" -verbose down

migrateforce:
	migrate -path database/migrations -database "sqlserver://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}" force $(VERSION)

dockerbuild:
	docker build -t find_peace:$(VERSION) .
