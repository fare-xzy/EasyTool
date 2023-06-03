package main

import (
	inLog "EasyToolServer/log"
	"net/http"
)

var (
	port = ":9527"
)

func main() {
	srv := &http.Server{Addr: port, Handler: InitHandler()}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			inLog.Errorf("端口被占用:%+v", err)
		}
	}()
	<-make(chan bool)
}
