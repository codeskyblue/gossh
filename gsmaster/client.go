package main

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/gcmurphy/getpass"
	"github.com/shavac/readline"
	"github.com/shxsun/gossh/gsmaster/rpc"
)

func Usage() {
	fmt.Printf("Usage: %s [user@]host\n\nCOPYRIGHT: sunshengxiang01@baidu.com\n", os.Args[0])
}

// split address user@hostname to user, hostname
func SplitUserHost(addr string) (username, host string, err error) {
	i := strings.Index(addr, "@")
	if i == -1 {
		u, er := user.Current()
		if er != nil {
			err = er
			return
		}
		username = u.Username
		host = addr
	} else {
		username, host = addr[:i], addr[i+1:]
	}
	return
}

func interactive(host string) (hostname, password string, err error) {
	// get hostname
	prompt := fmt.Sprintf("hostname[%s] > ", host)
	result := readline.ReadLine(&prompt)
	if result == nil || *result == "" { // EOF
		hostname = host
	} else {
		fmt.Println(*result)
		hostname = *result
	}
	// get password
	if password, err = getpass.GetPass(); err != nil {
		return
	}
	if password == "" {
		err = errors.New("password empty")
	}
	return
}

type HostInfo struct {
	Hostname string
	Username string
	Password string
}

// FIXME: not finished yet
func (hi *HostInfo) CheckSshConnection() (err error) {
	fmt.Printf("sshpass -p %s ssh -l %s %s\n", hi.Password, hi.Username, hi.Hostname)
	return nil
}

// FIXME: ...
func (hi *HostInfo) Connect() {
	fmt.Println("connecting ...")
}

func handleClient(client *rpc.GsClient) (err error) {
	if flag.NArg() < 1 {
		Usage()
		os.Exit(1)
	}
	addr := flag.Arg(0)
	var password string
	user, host, err := SplitUserHost(addr)
	if err != nil {
		return
	}

	r, err := client.LookHost(host, user)
	if err != nil {
		panic(err)
	}
	// parse server info
	if r.ErrorA1 != "" {
		fmt.Println(r.ErrorA1)
		host, password, err = interactive(host)
		if err != nil {
			return
		}
	}
	fmt.Println(user, host, password, err)
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
