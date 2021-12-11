hello:
	echo "hello"

run:
	@go run main.go

build:
	@go build -ldflags="-X 'github.com/BigGroupProgramming/discord-clone/app.Version=${VERSION}' -X 'github.com/BigGroupProgramming/discord-clone/app.Release=${RELEASE}' -X 'github.com/BigGroupProgramming/discord-clone/app.BuildTime=${NOW}'" -o build
.PHONY: build
