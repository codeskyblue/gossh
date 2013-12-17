package main

import (
	"flag"
	"fmt"
	"log"

	//"github.com/coocood/jas"
	"github.com/shxsun/gossh/gsmaster/models"
)

//type V1 struct{}

//func (*V1) PostLogin(ctx *jas.Context) {
//	user := ctx.RequireString("username")
//	pass := ctx.RequireString("password")
//	fmt.Println(user, pass)
//	ctx.Data = "success"
//}

//func (*V1) GetList(ctx *jas.Context) {
//}

func Test() {
	host := new(models.Host)
	host.Hostname = "xyzdas"
	host.Alias = "abc"
	n, err := models.Engine.InsertOne(host)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
}

var (
	server = flag.Bool("server", false, "run as a server")
	token  = flag.String("token", "abcdefg", "auth token for check identity")
)

func main() {
	flag.Parse()
	if *server {
		fmt.Println("goshh master ...")
		models.CreateDb()
	}

	Test()
	//router := jas.NewRouter(new(V1))
	//fmt.Println(router.HandledPaths(true))
	//http.Handle(router.BasePath, router)
	//err := http.ListenAndServe(":8033", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//hu := new(models.HostUser)
}
