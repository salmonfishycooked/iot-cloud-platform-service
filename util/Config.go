package util

import (
	"bufio"
	"encoding/json"
	"os"
)

// Config 总配置项
type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Database DatabaseConfig `json:"database"`
}

// DatabaseConfig 数据库配置项
type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

var cfg *Config // 配置项实例

const path = "./config/app.json" // app 配置项路径

// ParseConfig
// @Description: 获取解析出来的 app 配置项结构体（单例模式）
// @return *Config app 配置项结构体
func ParseConfig() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
