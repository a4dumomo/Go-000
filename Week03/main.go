package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		return Signal(ctx)
	})

	eg.Go(func() error {
		return Server(ctx, ":8081")
	})

	if err := eg.Wait(); err != nil {
		log.Println("err:", err)
	}
	log.Println("service closed success")
}

//router
func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World:" + request.RemoteAddr))
	})
	return mux
}

func Signal(ctx context.Context) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-quit:
		return errors.New("receive signal stop")
	case <-ctx.Done():
		log.Println("signal listen sotop")
		return nil
	}
}

//server
func Server(ctx context.Context, addr string) error {
	server := http.Server{
		Addr:    addr,
		Handler: Mux(),
	}
	stop := make(chan error, 1)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("shutdown err")
			}
		}()
		if err := server.ListenAndServe(); err != nil {
			log.Println("start service fail:", err)
			stop <- err
		}
	}()

	select {
	case <-ctx.Done():
		ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx2); err != nil {
			return err
		}
	case err := <-stop:
		return err
	}
	return nil
}
