server:
	go run main.go
proxy:
	go run proxy/proxy.go
gen:
	protoc --go_out=. demo/*.proto
	protoc --go-grpc_out=. demo/*.proto
	protoc -I . --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true demo/demo.proto
.PHONY: all server client gen proxy clean 