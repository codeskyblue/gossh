package main

import "github.com/shxsun/gossh/gsmaster/rpc"

type GsHandler struct{}

func NewGsHandler() *GsHandler {
	return &GsHandler{}
}

func (p *GsHandler) LookHost(hostname string, username string) (r *rpc.Data, err error) {
	r = &rpc.Data{}
	r.Hostname = "abcdefg"
	return r, nil
}

func (p *GsHandler) SyncHost(data *rpc.Data) (r bool, err error) {
	return
}
