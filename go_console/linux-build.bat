set GOOS=linux
set CGO_ENABLED=0
set GO111MODULE=on

del bin/console
go build -o bin/console services/console/main.go
REM pause