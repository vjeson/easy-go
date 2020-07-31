package main

import (
	"demo/conf"
	"demo/db"
	"demo/router"
)

func main()  {
	defer db.Db.Close()
	router := router.InitRouter()
	router.Run(conf.Conf.Server.Address)

	//util.GenTableToStruct()
}
