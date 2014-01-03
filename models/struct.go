// struct.go
package models

/* Structure */

type Record struct {
	Hostname  string `xorm:"unique(u)"`
	User      string `xorm:"unique(u)"`
	Pass      string
	ViewCount int64 `xorm:"default 0"`
}

//type User struct {
//	Id       int64 `xorm:"pk autoincr"`
//	Username string
//	Password string
//}

//type Host struct {
//	Id        int64 `xorm:"pk autoincr"`
//	Alias     string
//	Hostname  string `xorm:"unique"`
//	ViewCount int64  `xorm:"default 0"`
//}

//type HostUser struct {
//	Host Host `xorm:"host_id"`
//	User User `xorm:"user_id"`
//}

//type Statistic struct {
//	Hostname  string
//	ViewCount int64
//}
