@echo "开始编译go_console"
set GOOS=linux
set CGO_ENABLED=0
set GO111MODULE=auto
cd go_console
del ../docker_console/go_console/go_console
go build -o ../docker_console/go_console/go_console ./services/console/main.go

cd ../vue_console
npm run build
REM pause