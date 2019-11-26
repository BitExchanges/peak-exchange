package utils

import (
	"peak-exchange/config"
	"testing"
)

func TestNewEnv(t *testing.T) {
	config.InitEnv()
	configEnv := getDatabaseConfig()
	if configEnv == nil {
		t.Error("加载数据库文件失败")
	}

}
