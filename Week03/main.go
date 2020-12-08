package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	quit := make(chan os.Signal, 1) //service stop signal
	done := make(chan struct{}, 1)  //service stop finsh
	eg, _ := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		return server(quit, done)
	})

	eg.Go(func() error {
		signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		<-done
		return nil
	})

	err := eg.Wait()
	if err != nil {
		log.Fatalf("Service monitoring failed:%v", err)
	}
	log.Println("The service has been exited")
}

//server
func server(quit <-chan os.Signal, done chan<- struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Helllo World!"))
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Shutdown Error:%v\n", err)
			}
		}()
		<-quit
		//5s to finish all service
		ctx, cancal := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancal()
		if err := server.Shutdown(ctx); err != nil {
			log.Println("Service shutdown failed,err:", err)
		}
		done <- struct{}{}
	}()
	return server.ListenAndServe()
}
