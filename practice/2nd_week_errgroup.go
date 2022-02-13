package practice

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

//直接go函数 无法处理goroutine的返回数据
//errgroup主要用于对goroutine返回数据进行处理error
func main() {
	ctx := context.Background()
	eg, _ := errgroup.WithContext(ctx)

	for i := 0; i < 3; i++ {
		num := i
		eg.Go(func() error {
			//do sth...
			time.Sleep(time.Second * time.Duration(num*2))
			fmt.Printf("goroutine %d done!\n", num)
			// if num == 0 {
			// 	return errors.New("into 0")
			// }
			if num == 1 {
				return errors.New("into 1")
			}
			if num == 2 {
				return errors.New("into 2")
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Println("eg error: ", err)
	} else {
		fmt.Println("all goroutine done!")
	}
}
