package models

import (
	"errors"
	"fmt"
)

func (r *Record) Add() (err error) {
	_, err = Engine.InsertOne(r)
	return
}

func (r *Record) Sync() (err error) {
	_, err = GetRecord(r.Hostname, r.User)
	if err != nil {
		return r.Add()
	}
	_, err = Engine.Update(r)
	return
}

func GetRecord(host, user string) (r *Record, err error) {
	r = new(Record)
	fmt.Println("DEBUG", host, user)
	ok, err := Engine.Where("hostname=? AND user=?", host, user).Get(r) //r.Hostname, r.User).Get(r)
	if !ok {
		return nil, errors.New("record not found")
	}
	return r, err
}

func LookHost(alias string) (hostname, password string, err error) {
	err = errors.New("match too many hosts")
	return
}
