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

func AddRecord(r *Record) (err error) {
	_, err = Engine.InsertOne(r)
	return
}

func GetRecord(host, user string) (r *Record, err error) {
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
	oldr, err := GetRecord(r.Hostname, r.User)
	log.Println(oldr, err)
	if oldr != nil && err == nil {
		_, err = Engine.Update(r)
		return
	} else {
		err = AddRecord(r)
		return
	}
}

//func LookHost(alias string) []

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

//type HostUser struct {
//	HostId int64 `xorm:"unique(m)"`
//	UserId int64 `xorm:"unique(m)"`
//}

/*
type User struct {
	Id                int64 `xorm:"pk"`
	Username          string
	HashLoginPassword string // login password
	ViewCount         int64
	BornTime          time.Time
	Type              string
	AesInfoPassword   string // passwd to crypt host:user,pass crypted by login password
}

type Group struct {
	UserId      int64 `xorm:"unique(g)"`
	GroupUserId int64 `xorm:"unique(g)"`
}

type HostUser struct {
	Id          int64 `xorm:"pk"`
	UserId      int64
	Username    string
	AesPassword string // crypted password by user
}

type Host struct {
	Hostname   string `xorm:"unique(uh)"`
	HostUserId int64  `xorm:"unique(uh)"`
	ViewCount  int64
}

// AesPassword: temp stored host:user,pass crypted too. but passwd store by user as token.
type VirtPassword struct {
	Id          int64 `xorm:"pk"`
	UserId      int64
	AesPassword string
}
*/
