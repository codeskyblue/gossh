package main

import (
	"flag"
	"fmt"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/shxsun/gossh/gsmaster/models"
)

//func Test() {
//	host := new(models.Host)
//	host.Hostname = "xyzdas"
//	host.Alias = "abc"
//	n, err := models.Engine.InsertOne(host)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println(n)
//}

var (
	server   = flag.Bool("server", false, "run as a server")
	token    = flag.String("token", "abcdefg", "auth token for check identity")
	addr     = flag.String("addr", "localhost:6523", "server address")
	secure   = flag.Bool("secure", false, "enable secure")
	protocol = "binary"
	buffered = false
	framed   = true
)

func main() {
	flag.Parse()
	if *server {
		fmt.Println("goshh master ...")
		models.CreateDb()
	}

	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json": // json is working well
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		//Usage()
		os.Exit(1)
	}

	var transportFactory thrift.TTransportFactory
	if buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	if *server {
		if err := runServer(transportFactory, protocolFactory, *addr, *secure); err != nil {
			fmt.Println("error running server:", err)
		}
	} else {
		if err := runClient(transportFactory, protocolFactory, *addr, *secure); err != nil {
			fmt.Println("error running client:", err)
		}
	}
}
