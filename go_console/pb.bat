protoc.exe --go_out=./pb -I./pb/proto  ./pb/proto/comm.proto
protoc.exe --go_out=./pb -I./pb/proto  ./pb/proto/db_host.proto
protoc.exe --go_out=./pb -I./pb/proto  ./pb/proto/db_user.proto
protoc.exe --go_out=./pb -I./pb/proto  ./pb/proto/msg_api.proto
protoc.exe --go_out=./pb -I./pb/proto  ./pb/proto/msg_console.proto
protoc.exe --go_out=./pb -I./pb/proto  ./pb/proto/msg_user.proto
pause