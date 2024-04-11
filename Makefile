
# Define the name of our executable here
EXECUTABLE = bin/app

all: $(EXECUTABLE) test

$(EXECUTABLE): cmd/app/main.go
	go build -o ./$@ $^

run: all
	./$(EXECUTABLE)

test:
	go test ./...

clean:
	rm -f $(EXECUTABLE)
