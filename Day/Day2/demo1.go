package Day2

import (
	"bufio"
	"fmt"
	"net"
)

// TCP server端
// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)

		conn.Write([]byte(recvStr)) // 发送数据
	}
}

// 网络编程
func Test1() {
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn 失败")
			continue
		}
		go process(conn)
	}
}
