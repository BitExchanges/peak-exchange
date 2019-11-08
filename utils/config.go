package utils

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"strconv"
	"time"
)

type ConfigEnv struct {
	configFile *yaml.File
}

// 读取数据库配置文件
func getDatabaseConfig() *ConfigEnv {
	return getConfig("database")
}

// 读取redis配置文件
func getRedisConfig() *ConfigEnv {
	return getConfig("redis")
}

func getConfig(name string) *ConfigEnv {
	filePath := fmt.Sprintf("config/%s.yaml", name)
	return NewEnv(filePath)
}

func NewEnv(configFile string) *ConfigEnv {
	env := &ConfigEnv{configFile: yaml.ConfigFile(configFile)}
	if env.configFile == nil {
		panic("配置文件打开失败:" + configFile)
	}
	return env
}

func (env *ConfigEnv) Get(spec, defaultValue string) string {
	value, err := env.configFile.Get(spec)
	if err != nil {
		value = defaultValue
	}
	return value
}

func (env *ConfigEnv) GetInt(spec string, defaultValue int) int {
	str := env.Get(spec, "")
	if str == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Panic("数据类型转换错误: int: ", spec, str)
	}
	return val
}

func (env *ConfigEnv) GetDuration(spec string, defaultValue string) time.Duration {
	str := env.Get(spec, "")
	if str == "" {
		str = defaultValue
	}
	duration, err := time.ParseDuration(str)
	if err != nil {
		log.Panic("数据类型转换错误: duration:", spec, str)
	}
	return duration
}
