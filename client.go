package main

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/gcmurphy/getpass"
	"github.com/shxsun/gossh/rpc"
)

var ErrPassword = errors.New("password error")

func Usage() {
	fmt.Printf("Usage: %s [user@]host [args...]\n\nCOPYRIGHT: sunshengxiang01@baidu.com\n", os.Args[0])
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
		return
	}
	username, host = addr[:i], addr[i+1:]
	return
}

type HostInfo struct {
	Hostname  string
	Username  string
	Password  string
	RpcClient *rpc.GsClient
}

// FIXME: not finished yet
func ReadString(prompt string) string {
	fmt.Print(prompt)
	return ""
}

func (hi *HostInfo) SyncToServer() (err error) {
	data := rpc.NewData()
	data.Hostname = hi.Hostname
	data.Username = hi.Username
	data.Password = hi.Password
	var ok bool
	ok, err = hi.RpcClient.SyncHost(data)
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("unexpected from server error")
	}
	return
}

func (hi *HostInfo) Interactive() (err error) {
	// get hostname
	prompt := fmt.Sprintf("hostname[%s@%s] > ", hi.Username, hi.Hostname)
	hostname := ReadString(prompt)
	if hostname != "" {
		hi.Hostname = hostname
	}
	// get password
	if hi.Password, err = getpass.GetPass(); err != nil {
		return
	}
	if hi.Password == "" {
		err = errors.New("password empty")
	}
	return
}

// used for sshpass
func (hi *HostInfo) GenSshpassArgs(args []string) []string {
	return append([]string{"-e", "ssh", "-l", hi.Username, hi.Hostname}, args...)
}

func (hi *HostInfo) CheckSshConnection() (err error) {
	Debugf("checking ...")
	cmd := exec.Command("sshpass", hi.GenSshpassArgs([]string{"echo", "1"})...)
	cmd.Env = []string{"SSHPASS=" + hi.Password}
	out, err := cmd.Output()
	if err != nil {
		switch err.Error() {
		case "exit status 5":
			return ErrPassword
		case "exit status 6":
			return fmt.Errorf("host: [%s] connected failed", hi.Hostname)
		}
		return
	}
	if string(out) != "1\n" {
		return ErrPassword
	}
	return nil
}

func (hi *HostInfo) Connect(args []string) (err error) {
	Debugf("connecting to %s ...\n", hi.Hostname)
	argv := hi.GenSshpassArgs(args)
	argv = append([]string{hi.Hostname}, argv...)
	//Debugf("%v\n", argv)
	path, err := exec.LookPath("sshpass")
	if err != nil {
		return
	}
	envs := []string{"SSHPASS=" + hi.Password, "TERM=screen"}
	return syscall.Exec(path, argv, envs)
}

func handleClient(client *rpc.GsClient) (err error) {
	if flag.NArg() < 1 {
		Usage()
		os.Exit(1)
	}
	addr := flag.Arg(0)
	user, host, err := SplitUserHost(addr)
	if err != nil {
		return
	}

	r, err := client.LookHost(host, user)
	if err != nil {
		panic(err)
	}
	// parse server info
	var hi = &HostInfo{
		Hostname:  host,
		Username:  user,
		Password:  "",
		RpcClient: client,
	}
	if r.ErrorA1 != "" {
		fmt.Println(r.ErrorA1)
		err = hi.Interactive()
		if err != nil {
			return
		}
	} else {
		hi.Password = r.Password
	}
	// start connecting checking
	for {
		if err = hi.CheckSshConnection(); err == nil {
			break
		}
		switch err {
		case ErrPassword:
			hi.Interactive()
		default:
			return
		}
	}
	if err = hi.SyncToServer(); err != nil {
		return
	}
	return hi.Connect(flag.Args()[1:])
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
