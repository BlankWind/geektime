package practice

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":8081"}
	http.HandleFunc("/", SimpleServer)
	fmt.Println("http server start!")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("server start error: ", err)
	}
}

func SimpleServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Geektime!\n")
}
