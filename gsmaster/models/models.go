package models

import (
	"errors"
	"log"
)

type Record struct {
	Hostname  string `xorm:"unique(u)"`
	User      string `xorm:"unique(u)"`
	Pass      string
	ViewCount int64 `xorm:"default 0"`
}

func (r *Record) Add() (err error) {
	_, err = Engine.InsertOne(r)
	return
}

func (r *Record) Get() (err error) {
	r.Pass, r.ViewCount = "", 0
	ok, err := Engine.Get(r)
	//ok, err := Engine.Where("hostname=? AND user=?", r.Hostname, r.User).Get(r)
	if !ok {
		return errors.New("record not found")
	}
	return err
}

func (r *Record) Sync() (err error) {
	err = r.Get()
	if err != nil {
		return r.Add()
	}
	_, err = Engine.Update(r)
	return
}

func addRecord(r *Record) (err error) {
	_, err = Engine.InsertOne(r)
	return
}

func getRecord(host, user string) (r *Record, err error) {
	r = new(Record)
	ok, err := Engine.Where("hostname=? AND user=?", host, user).Get(r)
	if !ok {
		return nil, nil
	}
	return
}

func DelRecord(host, user string) error {
	r := new(Record)
	r.Hostname, r.User = host, user
	_, err := Engine.Delete(r)
	return err
}

func SyncRecord(r *Record) (err error) {
	oldr, err := getRecord(r.Hostname, r.User)
	log.Println(oldr, err)
	if oldr != nil && err == nil {
		_, err = Engine.Update(r)
		return
	} else {
		err = addRecord(r)
		return
	}
}

func LookHost(alias string) (hostname, password string, err error) {
	err = errors.New("match too many hosts")
	return
}

//type Statistic struct {
//	Hostname  string
//	ViewCount int64
//}

type User struct {
	Id       int64 `xorm:"pk autoincr"`
	Username string
	Password string
}

type Host struct {
	Id        int64 `xorm:"pk autoincr"`
	Alias     string
	Hostname  string `xorm:"unique"`
	ViewCount int64  `xorm:"default 0"`
}

type HostUser struct {
	Host Host `xorm:"host_id"`
	User User `xorm:"user_id"`
}
