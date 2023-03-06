package tcp

import (
	"fmt"
	"net"
	"time"
)

// TCP 业务逻辑
func process(conn net.Conn) {
	// 函数执行完之后关闭连接
	defer conn.Close()
	for {
		var buf [128]byte
		// 将tcp连接读取到的数据读取到byte数组中, 返回读取到的byte的数目
		// 对连接设置 读 超时时间
		conn.SetReadDeadline(time.Now().Add(time.Duration(TIMEOUT) * time.Second))
		n, err := conn.Read(buf[:])
		if err != nil {
			// 从客户端读取数据的过程中发生错误（客户端断开服务器）
			fmt.Println("设备 (", conn.RemoteAddr(), ") 已断开服务器。")
			break
		}
		recStr := string(buf[:n])
		fmt.Println("收到客户端 (", conn.RemoteAddr(), ")", "发来的数据：", recStr)

		// 由于是tcp连接所以双方都可以发送数据, 下面接收服务端发送的数据这样客户端也可以收到对应的数据
		//inputReader := bufio.NewReader(os.Stdin)
		//s, _ := inputReader.ReadString('\n')
		//t := strings.Trim(s, "收不收得到")
		// 向当前建立的tcp连接发送数据, 客户端就可以收到服务端发送的数据
		//conn.Write([]byte("can you receive my message?"))
		//go sendMessage(conn)
	}
}

func sendMessage(conn net.Conn) {
	for {
		conn.Write([]byte("\"apitag\":\"Vehicle_condition1\", \"data\":1"))
	}
}
