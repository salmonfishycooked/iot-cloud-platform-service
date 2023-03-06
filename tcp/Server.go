package tcp

import (
	"fmt"
	"net"
)

const (
	IP      = "192.168.3.55" // 监听IP
	PORT    = "20000"        // 监听端口
	TIMEOUT = 5              // 请求超时时间（单位：秒）
)

// StartListenTCP 开始监听 TCP 连接
func StartListenTCP() {
	listen, err := net.Listen("tcp", IP+":"+PORT) // 监听端口
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	fmt.Println("已开启 TCP 监听（在 " + IP + ":" + PORT + " 上）")
	for {
		conn, err := listen.Accept() // 建立连接
		fmt.Println("检测到新的连接: ", conn.RemoteAddr())
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		// 对于每一个建立的 TCP 连接开启新线程去处理
		go process(conn)
	}
}
