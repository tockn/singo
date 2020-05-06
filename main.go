package main

import (
	"flag"
	"fmt"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var (
		withExample = flag.Bool("example", false, "serve with video chat system example")
		addrFlag    = flag.String("addr", "0.0.0.0", "addr")
		portFlag    = flag.Int("port", 5000, "port")
	)
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *addrFlag, *portFlag)

	if *withExample {
		return serveWithExample(addr)
	}

	return serve(addr)
}
