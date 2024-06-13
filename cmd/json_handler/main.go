package main

import "flag"

var (
	addr          = flag.String("port", "8080", "port for the json handler")
	protoEndpoint = flag.String("proto-endpoint", "localhost:8081", "proto handler server endpoint")
)

func main() {
	flag.Parse()
	InitializeServer(*addr, *protoEndpoint)
}
