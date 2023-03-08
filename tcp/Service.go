package tcp

import (
	"encoding/json"
	"errors"
	"fmt"
	"iot_backend/dao"
	"iot_backend/model"
)

// TCP 业务逻辑
func process(conn Connection) {
	canAccess := authDevice(&conn) // 对设备进行鉴权（等待设备发送tag）
	// 鉴权不通过直接踢掉连接
	if !canAccess {
		conn.conn.Write([]byte("鉴权未通过！不存在该设备或者该设备已经上线！"))
		conn.conn.Close()
		fmt.Println("连接", conn.conn.RemoteAddr(), "已被踢出（未通过鉴权）")
		return
	}
	//通过鉴权
	conn.conn.Write([]byte("鉴权通过！您的设备标识名: " + conn.DeviceTag))
	server.connections = append(server.connections, conn) // 加入连接数组里面，方便管理

	go keepAlive(conn) // 进行心跳检测
	for {
		recStr, err := readFromClient(conn.conn) // 从客户端读数据
		if err != nil {
			break
		}

		if SHOW_DATA {
			fmt.Println("设备", conn.DeviceTag, "发来的数据：", recStr) // 设备发来的数据
		}
		go createHistorySensorData(recStr, conn.DeviceTag) // 更新传感器的值
	}
}

// SendOrder
// @Description: 发送执行器传来的命令
// @param order 命令数据
// @param deviceTag 需要发送的设备tag
// @return error
func SendOrder(order string, deviceTag string) error {
	for _, val := range server.connections {
		// 找到对应设备，发出命令
		if deviceTag == val.DeviceTag {
			_, err := val.conn.Write([]byte(order))
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("error")
}

// createHistorySensorData
// @Description: 更新传感器的值
// @param str
// @return error
func createHistorySensorData(str string, deviceTag string) {
	data := SensorData{}
	if err := json.Unmarshal([]byte(str), &data); err == nil {
		err, count := dao.UpdateSensor(deviceTag, data.SensorTag, data.Value)
		// 如果没有找到该传感器
		if count == 0 {
			return
		}
		// 在历史传感数据里面新增一条记录
		historySensor := model.HistorySensorData{
			DeviceTag: deviceTag,
			SensorTag: data.SensorTag,
			Value:     data.Value,
		}
		dao.CreateData(&historySensor)
		if err == nil && SHOW_CREATE {
			fmt.Println("(√) 收到 " + deviceTag + " 的传感器 " + data.SensorTag + " 的传感值: " + data.Value)
		}
	}
}
