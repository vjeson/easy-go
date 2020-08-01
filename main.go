package main

import (
	"context"
	"demo/conf"
	"demo/db"
	"demo/router"
	"demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main()  {
	defer db.Db.Close()
	gin.SetMode(conf.Conf.Server.Model)
	router := router.InitRouter()
	//router.Run(conf.Conf.Server.Address)

	srv := &http.Server{
		Addr:    conf.Conf.Server.Address,
		Handler: router,
	}

	log := util.Log()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
