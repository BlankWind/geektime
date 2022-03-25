package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

/*
goim 协议结构
4bytes PacketLen 包长度
2bytes HeaderLen 头长度
2bytes Version 协议版本号
4bytes Operation 业务操作码
4bytes Sequence 序列号
PacketLen-HeaderLen Body 业务数据
*/

const (
	PackageLength   = 4
	HeaderLength    = 2
	ProtocolVersion = 2
	OperationLength = 4
	SequenceID      = 4
	AllHeaderLength = PackageLength + HeaderLength + ProtocolVersion + OperationLength + SequenceID

	protocolVersion = 2
	operation       = 3
	sequence        = 9
)

func encoder(body string) []byte {
	packetLen := len(body) + AllHeaderLength
	res := make([]byte, packetLen)

	binary.BigEndian.PutUint32(res[:4], uint32(packetLen))
	binary.BigEndian.PutUint16(res[4:6], uint16(AllHeaderLength))
	binary.BigEndian.PutUint16(res[6:8], uint16(protocolVersion))
	binary.BigEndian.PutUint32(res[8:12], uint32(operation))
	binary.BigEndian.PutUint32(res[12:16], uint32(sequence))

	byteBody := []byte(body)
	copy(res[16:], byteBody)

	return res
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	connHandler(conn)
}

func connHandler(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	fmt.Println("Please input data...")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ReadString err:", err)
			return
		}
		//编码数据
		c.Write(encoder(input))
		_, err = c.Read(buf)
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		fmt.Printf("server response:%v", string(buf))
	}
}
