package models

import (
	"errors"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/lib/pq" //当某时间字段表现为0001-01-01 07:36:42+07:36:42形式的时候 会读不出数据
	//_ "github.com/bylevel/pq"
	"github.com/lunny/xorm"
	"log"
)

var (
	Engine *xorm.Engine
)

const (
	DbName         = "./data/sqlite.db"
	DbUser         = "root"
	mysqlDriver    = "mymysql"
	mysqlDrvformat = "%v/%v/"
	pgDriver       = "postgres"
	pgDrvFormat    = "user=%v dbname=%v sslmode=disable"
	dbtype         = "sqlite"
)

func init() {
	_, err := SetEngine()
	if err != nil {
		log.Println(err)
	}
}

func XConDb() (*xorm.Engine, error) {
	switch {
	case dbtype == "sqlite":
		return xorm.NewEngine("sqlite3", DbName)

	case dbtype == "mysql":
		return xorm.NewEngine("mysql", "user=mysql password=jn!@#9^&* dbname=mysql")

	case dbtype == "pgsql":
		return xorm.NewEngine("postgres", "user=postgres password=jn!@#$%^&* dbname=pgsql sslmode=disable")
	}
	return nil, errors.New("尚未设定数据库连接")
}

func SetEngine() (*xorm.Engine, error) {
	var err error
	Engine, err = XConDb()
	//Engine.Mapper = xorm.SameMapper{}
	//Engine.SetMaxConns(5)
	//Engine.ShowSQL = true

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	Engine.SetDefaultCacher(cacher)

	return Engine, err
}

func CreateDb() {
	err := Engine.Sync(new(User), new(Group), new(HostUser), new(Host), new(VirtPassword))
	if err != nil {
		log.Fatal(err)
	}
}
