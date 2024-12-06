# Makefile
# Makefile is used to make runnig commands easy

# Define the command as a variable
# Once you use .PHONY for server, make will always execute the go run command, even if there are no file changes.
.PHONY: server
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
	go build -o ../bin/hms main.go && cp data/hms.db ../bin/data && cp env.example ../bin && cp .env ../bin
