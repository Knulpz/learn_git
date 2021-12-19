package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

//errgroup实现http_server的启动和关闭;linux信号的注册和处理；且保证一个goroutine退出，全部退出。
//实现方案：
//- http server的启动和关闭
//- 使用chan实现对linux signal中断的注册和处理
//- 使用errgroup实现级联注销

func main() {
	ctx := context.Background()
	// 定义 withCancel -> cancel() 方法 去取消下游的 Context
	ctx, cancel := context.WithCancel(ctx)
	// 使用 errgroup 进行 goroutine 取消
	g, errCtx := errgroup.WithContext(ctx)
	srv := &http.Server{Addr: ":8027"}

	//启动http server
	g.Go(func() error {
		return HttpServer(srv)
	})

	//关闭http server
	g.Go(func() error {
		<-errCtx.Done() //阻塞。cancel、timeout、deadline 都可能导致 Done 被 close
		fmt.Println("http server stop")
		return srv.Shutdown(errCtx) // 关闭 http server
	})

	//使用channel实现对signal的处理
	channel := make(chan os.Signal, 1) //buffer为1
	signal.Notify(channel)

	g.Go(func() error {
		for {
			select {
			case <-errCtx.Done(): //cancel/timeout/deadline
				return errCtx.Err()
			case <-channel: //kill - 9或其他signal
				cancel()
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("finished")
}

func HttpServer(srv *http.Server) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi Voyager")
	})
	fmt.Println("http server start")
	err := srv.ListenAndServe()
	return err
}
