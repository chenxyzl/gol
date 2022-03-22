#mac
go build -o CodeGenerator main.go

#win
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o CodeGenerator.exe main.go 

