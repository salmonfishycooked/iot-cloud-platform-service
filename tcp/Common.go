package tcp

import (
	"fmt"
	"net"
)

// ReadFromClient
// @Description: 从客户端读数据
// @param conn
// @return string 读取到的字符串
// @return error
func readFromClient(conn net.Conn) (string, error) {
	var buf [256]byte
	// 将tcp连接读取到的数据读取到byte数组中, 返回读取到的byte的数目
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	recStr := string(buf[:n])

	return recStr, nil
}
