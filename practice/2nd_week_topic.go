package practice

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

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
		srv.Shutdown(errCtx) //shutdown内部有ctx.done()，所以这一步就返回了error
		return errors.New("shutdown error!")

	})

	c := make(chan os.Signal, 1)
	signal.Notify(c)

	eg.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				fmt.Println("err Stop!")
				time.Sleep(time.Second * time.Duration(2))
				return errors.New("done error!")
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
