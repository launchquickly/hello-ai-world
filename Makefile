
# Define the name of our executables here
EXECUTABLE_APP = bin/app
EXECUTABLE_HTTP = bin/http

all: $(EXECUTABLE_APP) $(EXECUTABLE_HTTP) test

$(EXECUTABLE_APP): cmd/app/main.go
	go build -o ./$@ $^

$(EXECUTABLE_HTTP): cmd/http/main.go
	go build -o ./$@ $^

run: all
	./$(EXECUTABLE_APP)

test:
	go test ./...

clean:
	rm -f $(EXECUTABLE_APP) $(EXECUTABLE_HTTP)

setup:
	go mod download