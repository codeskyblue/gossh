package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coocood/jas"
	"github.com/shxsun/gossh/gsmaster/models"
)

type V1 struct{}

func (*V1) PostLogin(ctx *jas.Context) {
	user := ctx.RequireString("username")
	pass := ctx.RequireString("password")
	fmt.Println(user, pass)
	ctx.Data = "success"
}

func (*V1) GetList(ctx *jas.Context) {
}

func main() {
	fmt.Println("gsmaster ...")
	models.CreateDb()
	router := jas.NewRouter(new(V1))
	fmt.Println(router.HandledPaths(true))
	http.Handle(router.BasePath, router)
	err := http.ListenAndServe(":8033", nil)
	if err != nil {
		log.Fatal(err)
	}
}
