package practice

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	a := <-c
	fmt.Println("Go signal:", a)
}
