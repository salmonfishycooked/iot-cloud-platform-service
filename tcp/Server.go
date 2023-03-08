package tcp

import (
	"fmt"
	"net"
)

const (
	PORT        = "20000" // 监听端口
	TIMEOUT     = 5       // 请求超时时间（单位：秒）
	SHOW_DATA   = true    // 是否显示设备发来的数据
	SHOW_CREATE = true    // 是否显示成功插入数据库的反馈信息
)

// Server TCP服务器实例
type Server struct {
	listener    net.Listener
	connections []Connection
}

// Connection 每个连接的实例
type Connection struct {
	conn      net.Conn
	DeviceTag string // 该参数能唯一标识设备
}

var server Server // TCP服务器实例

// StartListenTCP 开始监听 TCP 连接
func StartListenTCP() {
	initTcpServer() // 初始化TCP服务器
	for {
		conn, err := server.listener.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		fmt.Println("检测到新的连接: ", conn.RemoteAddr())
		connection := Connection{
			conn:      conn,
			DeviceTag: "",
		}

		go process(connection) // 对于每一个建立的 TCP 连接开启新线程去处理
	}
}

// initTcpServer
// @Description: 初始化TCP服务器
func initTcpServer() {
	listen, err := net.Listen("tcp", ":"+PORT) // 监听端口
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	server.listener = listen

	fmt.Println("已开启 TCP 监听（在 " + "localhost:" + PORT + " 上）")
}
