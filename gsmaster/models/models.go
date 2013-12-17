package models

type User struct {
	Id       int64 `xorm:"pk"`
	Username string
	Password string
}

type Host struct {
	Id        int64 `xorm:"pk"`
	Alias     string
	Hostname  string
	ViewCount int64
}

type HostUser struct {
	HostId int64 `xorm:"unique(m)"`
	UserId int64 `xorm:"unique(m)"`
}

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
