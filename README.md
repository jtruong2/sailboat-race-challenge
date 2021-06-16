# Sailboat Race Challenge ⛵
This terminal-based tool takes in `finish times` of sailboat race competitors and returns the average time it took to complete the race in minutes.

### Setup
1. `git clone https://github.com/jtruong2/sailboat-race-challenge.git`
2. `cd` into project root
3. `go get` to install dependencies
### Run
1. `go run main.go` to start the program
2. Follow the on-screen prompts

### Testing
Run test with coverage
1. `go test ./... --coverprofile cover.out`
2. `go tool cover -func=cover.out`
### Dependencies
- [testify](https://github.com/stretchr/testifyhttps://github.com/stretchr/testify) - Used for test assertions