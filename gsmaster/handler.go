package main

import (
	"fmt"

	"github.com/shxsun/gossh/gsmaster/models"
	"github.com/shxsun/gossh/gsmaster/rpc"
)

type GsHandler struct{}

func NewGsHandler() *GsHandler {
	return &GsHandler{}
}

func (p *GsHandler) LookHost(host string, user string) (r *rpc.Data, err error) {
	r = &rpc.Data{}
	rd, err := models.GetRecord(host, user)
	if err != nil {
		r.ErrorA1 = err.Error()
		return r, nil
	}

	r.Hostname = rd.Hostname
	r.Password = rd.Pass
	r.Username = rd.User

	return r, nil
}

func (p *GsHandler) SyncHost(data *rpc.Data) (b bool, err error) {
	r := new(models.Record)
	r.Hostname = data.Hostname
	r.User = data.Username
	r.Pass = data.Password

	b = true
	if err := r.Sync(); err != nil {
		fmt.Println(err)
		b = false
	}
	return b, nil
}
