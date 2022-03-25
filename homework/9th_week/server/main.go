package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
)

func decoder(data []byte) {
	packetLen := binary.BigEndian.Uint32(data[:4])
	headerLen := binary.BigEndian.Uint16(data[4:6])
	version := binary.BigEndian.Uint16(data[6:8])
	operation := binary.BigEndian.Uint32(data[8:12])
	sequence := binary.BigEndian.Uint32(data[12:16])
	body := string(data[16:])

	//输出
	fmt.Printf("packet length:%v\n", packetLen)
	fmt.Printf("header length:%v\n", headerLen)
	fmt.Printf("version:%v\n", version)
	fmt.Printf("operation:%v\n", operation)
	fmt.Printf("sequence:%v\n", sequence)
	fmt.Printf("body:%v\n", body)
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	fmt.Println("Listening....")
	fmt.Println()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			return
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		decoder(line)
		fmt.Println("-----------------------------------------------")
		wr.WriteString("Receive Success!\n")
		wr.Flush()
	}
}
