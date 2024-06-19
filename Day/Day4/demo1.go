package Day4

import (
	"fmt"
	"net"
	"strconv"
)

// 服务端
type Server struct {
	IP   string
	Port int
}

// 工厂模式
func NewServer(IP string, Port int) *Server {
	server := &Server{
		IP:   IP,
		Port: Port,
	}
	return server
}

// 处理方法
func (this *Server) Handlerfunc(conn net.Conn) error {
	fmt.Println("连接成功")
	return nil
}

func (this *Server) Start() {
	listen, err := net.Listen("tcp", this.IP+":"+strconv.Itoa(this.Port))
	if err != nil {
		fmt.Printf("error listening: %s\n", err)
		return
	}
	defer listen.Close()

	for {
		// listen accpet
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("this accepting error : %s\n", err)
			continue
		}
		// handler
		go this.Handlerfunc(conn)
	}
}
