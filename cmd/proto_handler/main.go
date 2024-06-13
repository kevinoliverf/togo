package main

import "flag"

var (
	addr = flag.String("port", "8081", "port for the proto handler")
)

func main() {
	InitializeServer(*addr)
}
