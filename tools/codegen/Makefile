all: deps


deps:
	GOOS=windows go build -o CodeGenerator.exe main.go
	GOOS=linux go build -o CodeGenerator main.go
	GOOS=darwin go build -o CodeGenerator_mac main.go