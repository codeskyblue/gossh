package main

import (
	"flag"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/shxsun/gossh/models"
)

var (
	//token   = flag.String("token", "abcdefg", "auth token for check identity") // no needed now
	server  = flag.Bool("server", false, "run as a server")
	addr    = flag.String("addr", "localhost:6523", "server address")
	verbose = flag.Bool("v", true, "show verbose info")

	secure = false
)

func main() {
	flag.Parse()
	if *server {
		fmt.Println("goshh master ...")
		models.CreateDb()
	}

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

	if *server {
		if err := runServer(transportFactory, protocolFactory, *addr, secure); err != nil {
			fmt.Println("error running server:", err)
		}
	} else {
		if err := runClient(transportFactory, protocolFactory, *addr, secure); err != nil {
			fmt.Println("error running client:", err)
		}
	}
}
