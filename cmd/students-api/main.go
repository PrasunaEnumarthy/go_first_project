package main

import (
	"context"
	//"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PrasunaEnumarthy/GO/internal/config"
	"github.com/PrasunaEnumarthy/GO/internal/config/http/handlers/student"
)

//import "fmt"

func main() {
	//LOAD CONFIG
	cfg :=config.MustLoad()
	//db setup
	//setup router
	
	router:=http.NewServeMux()
	router.HandleFunc("/api/students",student.New())
	//setup server

	
	
	server:=http.Server{
		Addr: cfg.Add,
		Handler: router,
	}
	slog.Info("server started",slog.String("address",cfg.Add))
	done:=make(chan os.Signal,1)

	signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)
	//go routine ensures smooth shut down 
	go func(){
		err:=server.ListenAndServe()
		if err!=nil{
			log.Fatal("failed to start server")
		}
	}()
	<-done

	slog.Info("shutting down this sever")

	ctx, cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err:=server.Shutdown(ctx);err!=nil{
		slog.Error("failed to shutdown server",slog.String("error",err.Error()))
	}
	slog.Info("server shutdown successfully")
	
	


}