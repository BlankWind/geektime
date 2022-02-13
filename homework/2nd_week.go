/*
题目：
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/
package homework

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	ctx, cancle := context.WithCancel(ctx)
	eg, errCtx := errgroup.WithContext(ctx)

	srv := &http.Server{Addr: ":8081"}
	eg.Go(func() error {
		return UpServer(srv)
	})

	eg.Go(func() error {
		<-errCtx.Done()
		fmt.Println("Server Stop!")
		return srv.Shutdown(errCtx)

	})

	c := make(chan os.Signal, 1)
	signal.Notify(c)

	eg.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-c:
				cancle()
			}
		}
	})
	if err := eg.Wait(); err != nil {
		fmt.Println("Group Error: ", err)
	} else {
		fmt.Println("All group done!")
	}
}

func UpServer(srv *http.Server) error {
	http.HandleFunc("/", SimpleServer)
	fmt.Println("Server start!")
	return srv.ListenAndServe()
}

func SimpleServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Geektime!\n")
}
