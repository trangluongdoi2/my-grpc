# protoc -I greet/proto --go_out=greet/proto --go-grpc_out=greet/proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative greet/proto/dummy.proto
# protoc -I greet/proto --go_out=. --go_opt=module=github.com/trangluongdoi2/my-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/trangluongdoi2/my-grpc greet/proto/dummy.proto
BIN_DIR = bin
PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client
SERVER_BIN = ${SERVER_DIR}
CLIENT_BIN = ${CLIENT_DIR}
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

.PHONY: greet calculator all
project := greet calculator

all: $(project) ## Generate Pbs and build

$(project):
	protoc -I $@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. $@/${PROTO_DIR}/*.proto
	go build -o ${BIN_DIR}/$@/${SERVER_BIN} ./$@/${SERVER_DIR}
	go build -o ${BIN_DIR}/$@/${CLIENT_BIN} ./$@/${CLIENT_DIR}
	chmod +x ${BIN_DIR}/$@/${SERVER_BIN} ${BIN_DIR}/$@/${CLIENT_BIN}
	