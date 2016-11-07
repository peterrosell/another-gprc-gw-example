gen-rpc:
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
	./example-api/document.proto

gen-gw:
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:. \
	./example-api/document.proto

gen-swagger:
	protoc -I/usr/local/include -I. \
	 -I${GOPATH}/src \
	 -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	 --swagger_out=logtostderr=true:. \
	 ./example-api/document.proto

build-grpc-server: gen-rpc
	go build -o grpc-server \
	./grpc/main.go

build-gw-server: gen-gw
	go build -o gw-server \
	./gw/main.go


build-all: gen-rpc gen-gw gen-swagger build-grpc-server build-gw-server

