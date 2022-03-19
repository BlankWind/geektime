package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client
var ctx context.Context

func init() {
	ctx = context.Background()
	err := initClient(ctx)
	if err != nil {
		panic("connect err!")
	}
	fmt.Println("connected!")
}

func initClient(ctx context.Context) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "172.16.1.94:6379",
		Password: "",
		DB:       0,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {

	// setData(10000, "num1w_size10", genValueSize(10)) //used_memory:872320  used_memory:1803312  diff:930992
	// setData(50000, "num5w_size10", genValueSize(10)) //used_memory:872472  used_memory:5396680  diff:4524208
	// setData(100000, "num10w_size10", genValueSize(10)) //used_memory:872624  used_memory:9921200  diff:9048576

	// setData(10000, "num1w_size100", genValueSize(100)) //used_memory:872136  used_memory:2763208 diff:1891072
	// setData(50000, "num5w_size100", genValueSize(100)) //used_memory:872288  used_memory:10196576  diff:9324288
	// setData(100000, "num10w_size100", genValueSize(100)) //used_memory:872440  used_memory:19521016  diff:18648576

	// setData(10000, "num1w_size1k", genValueSize(1000)) //used_memory:872592  used_memory:11883584  diff:11010992
	// setData(50000, "num5w_size1k", genValueSize(1000)) //used_memory:872744  used_memory:55796952  diff:54924208
	setData(100000, "num10w_size1k", genValueSize(1000)) //used_memory:872896  used_memory:110721472  diff:109848576

	fmt.Println("Done!")

}

func setData(num int, key, value string) {
	var builder strings.Builder
	for i := 0; i < num; i++ {
		builder.WriteString(key)
		builder.WriteString("_")
		builder.WriteString(strconv.Itoa(i))
		k := builder.String()

		err := rdb.Set(ctx, k, value, 0).Err()
		if err != nil {
			fmt.Println("set failed, i:", i)
			builder.Reset()
			continue
		}
		builder.Reset()
	}
}

func genValueSize(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 'a'
	}
	return string(arr)
}
