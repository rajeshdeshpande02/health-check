package main

//https://github.com/eddycjy/go-gin-example/blob/master/routers/api/auth.go
import (
	"health-check/server"

	"github.com/golang/glog"
)

func main() {
	//utils.InitLogger()
	//flag.Parse()
	//defer glog.Flush()
	glog.Info("This is a log,", "level", "info")
	//	glog.Warning("This is a log,", "level", "warning")
	//	glog.Error("This is a log,", "level", "error")
	//glog.Fatal("This is a log,", "level", "fatal")
	server.Init()
}
