package main

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"strings"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/shxsun/gossh/gsmaster/rpc"
)

func handleClient(client *rpc.GsClient) (err error) {
	var host, user string
	if flag.NArg() < 1 {
		return errors.New("new at least one arguments")
	}
	host = flag.Arg(1)
	if i := strings.Index(host, "@"); i == -1 {
		fmt.Println("use default user:", "root")
		user = "root"
	}

	r, err := client.LookHost(host, user)
	if err != nil {
		return
	}
	fmt.Println("result=", r)
	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err

	}
	transport = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	return handleClient(rpc.NewGsClientFactory(transport, protocolFactory))
}
