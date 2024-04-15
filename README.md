# hello-ai-world

A simple Go program that greets the world using AI.

## Prerequisites

1. Install Go 1.22, if you haven't yet, ensure it is properly installed and configured for your system. You can verify 
   this by running `go version` in your terminal/command line. It should return a version of 1.22.x. If not, refer to 
   the official [Go installation guide](https://golang.org/doc/install).

## Setup

1. Clone the repository to your local machine with `git clone https://github.com/launchquickly/hello-ai-world.git`.
2. Navigate into the project directory in your terminal/command line with `cd hello-ai-world`.
3. Run `make setup` to install dependencies as listed in the 'go.mod' file. This ensures that all necessary packages are 
   installed for the project to run successfully.
4. To clean up compiled files, use `make clean`.

## Build

1. Run `make` to compile and create an executable file named 'app' and 'http'. This will generate a binary file that is 
   located inside the ./bin directory of your repository.

## Test

1. Run `make test` to execute all tests in the project.

## Run

1. Run `./bin/app` or `./bin/http` to execute the desired program. 
   1. `./bin/app` can also be passed a name as an argument, like so: `./bin/app "Paul"`.
   2. `./bin/http` can be accessed via a browser on http://localhost:8080/generateMessage passing in an argument as a 
       query string key, value pair

## Clean

1. To clean up compiled files, use `make clean`.

## Usage Information

The project uses Go language and doesn't require any software other than a working Go environment. It also includes 
tests, setup instructions, build, clean targets to ease development processes.
