package main

import (
	"errors"

	"github.com/KantaHasegawa/OAuth2.0_client_server/router"
)
func main() {
	r := router.NewRouter()
	if err := r.Run(":8081"); err != nil{
		panic(errors.New("can not start server"))
	}
}
