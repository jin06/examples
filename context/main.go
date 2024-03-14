package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	run5()
}

func run1() {
	// father
	go func() {
		// son
		go func() {
			for t := range time.Tick(time.Second) {
				fmt.Println("work", t)
			}
			fmt.Println("work done")
		}()
		<-time.Tick(2 * time.Second)
		fmt.Println("bye")
	}()
	select {}
}

func run2() {
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go func(ctx context.Context) {
			defer fmt.Println("work done")
			for {
				ticker := time.NewTicker(time.Second)
				select {
				case <-ctx.Done():
					return
				case t := <-ticker.C:
					fmt.Println("work", t)
				}
			}
		}(ctx)
		<-time.Tick(2 * time.Second)
		fmt.Println("bye")
	}()
	select {}
}

func run3() {
	ctxA, cancelA := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		fmt.Println("A start")
		defer fmt.Println("A stop")
		ctxB, _ := context.WithCancel(ctxA)
		go func(ctx context.Context) {
			fmt.Println("B start")
			defer fmt.Println("B stop")
			ctxC, _ := context.WithCancel(ctx)
			go func(ctx context.Context) {
				fmt.Println("C start")
				defer fmt.Println("C stop")
				<-ctx.Done()
			}(ctxC)
			<-ctx.Done()
		}(ctxB)
		<-ctx.Done()
	}(ctxA)
	<-time.Tick(time.Second)
	cancelA()
	<-time.Tick(time.Second)
}

func contextValue() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "request_id", "100")

	service := func(ctx context.Context) {
		id := ctx.Value("request_id")
		fmt.Println(id)
	}

	service(ctx)
}

// func run4() {
// ctx := context.Background()
// c, cancel := context.WithCancel(ctx)
// if cancel() c.Done()
// c2, cancel2 := context.WithDeadline(ctx, time.Now().Add(time.Second*10))
// context.WithTimeout()
// }

func run5() {
	srv := Server{
		closed: make(chan struct{}),
	}
	go func() {
		defer srv.Stop()
		go srv.Run()
		<-time.Tick(time.Second)
	}()
	<-srv.Stopped()
}

type Server struct {
	closed chan struct{}
}

func (s *Server) Run() {
	fmt.Println("run")
	<-s.closed
}

func (s *Server) Stop() {
	close(s.closed)
}

func (s *Server) Stopped() <-chan struct{} {
	return s.closed
}
