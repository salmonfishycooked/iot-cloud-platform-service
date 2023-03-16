package tcp

import (
	"fmt"
	"iot_backend/dao"
	"time"
)

const (
	AlivePeriod = 10 // 发送心跳检测周期（单位：s）
)

// keepAlive
// @Description: 检测连接是否还存在（心跳检测）
// @param conn 需要检测的已经建立的连接
func keepAlive(conn Connection) {
	fmt.Println("设备" + conn.DeviceTag + "已上线！")
	dao.UpdateDeviceStatus(conn.DeviceTag, true) // 更新设备在线状态
	for {
		//conn.conn.SetReadDeadline(time.Now().Add(time.Duration(TIMEOUT) * time.Second)) // 对连接设置 读 超时时间
		_, err := conn.conn.Write([]byte("a"))
		if err != nil {
			fmt.Println("设备" + conn.DeviceTag + "已离线！")
			break
		}
		time.Sleep(time.Duration(AlivePeriod) * time.Second)
	}
	conn.conn.Close()                             // 关闭连接
	dao.UpdateDeviceStatus(conn.DeviceTag, false) // 更新设备在线状态

	//将实例从数组中移出去
	i := 0
	for ; i < len(server.connections); i++ {
		if server.connections[i].DeviceTag == conn.DeviceTag {
			break
		}
	}
	server.connections = append(server.connections[:i], server.connections[i+1:]...) // 删除索引为i的元素
}
