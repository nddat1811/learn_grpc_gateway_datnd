server:
	go run main.go
proxy:
	go run proxy/proxy.go
.PHONY: all server client  proxy clean 