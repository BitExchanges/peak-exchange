package utils

import (
	"fmt"
	"github.com/streadway/amqp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

type Amqp struct {
	Connect struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"connect"`

	Exchange struct {
		Default  map[string]string `yaml:"default"`
		Matching map[string]string `yaml:"matching"`
		Trade    map[string]string `yaml:"trade"`
		Cancel   map[string]string `yaml:"cancel"`
		Fanout   map[string]string `yaml:"fanout"`
	} `yaml:"exchange"`
	Queue struct {
		Matching map[string]string `yaml:"matching"`
		Trade    map[string]string `yaml:"trade"`
		Cancel   map[string]string `yaml:"cancel"`
	} `yaml:"queue"`
}

var (
	AmqpGlobalConfig Amqp
	RabbitMqConnect  *amqp.Connection
)

// 初始化AMQP配置
func InitializeAmqpConfig() {
	path_str, _ := filepath.Abs("config/amqp.yml")
	content, err := ioutil.ReadFile(path_str)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = yaml.Unmarshal(content, &AmqpGlobalConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	InitializeAmqpConnection()
}

// 初始化AMQP连接
func InitializeAmqpConnection() {
	var err error
	RabbitMqConnect, err = amqp.Dial("amqp://" + AmqpGlobalConfig.Connect.Username + ":" + AmqpGlobalConfig.Connect.Password + "@" + AmqpGlobalConfig.Connect.Host + ":" + AmqpGlobalConfig.Connect.Port + "/")
	if err != nil {
		time.Sleep(5000)
		InitializeAmqpConnection()
		return
	}
	go func() {
		<-RabbitMqConnect.NotifyClose(make(chan *amqp.Error))
		InitializeAmqpConnection()
	}()
}

// 关闭AMQP连接
func CloseAmqpConnection() {
	RabbitMqConnect.Close()
}

// 获取AMQP连接
func GetRabbitMqConnect() *amqp.Connection {
	return RabbitMqConnect
}

func PublishMessageWithRouteKey(exchange, routeKey, contentType string, message *[]byte, arguments amqp.Table, deliveryMode uint8) error {
	channel, err := RabbitMqConnect.Channel()
	defer channel.Close()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	if err = channel.Publish(
		exchange, //交换机
		routeKey, //交换机与队列的路由key
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     contentType,
			ContentEncoding: "",
			DeliveryMode:    deliveryMode,
			Priority:        0,
			Body:            *message,
		},
	); err != nil {
		return fmt.Errorf("Queue publish: %s", err)
	}
	return nil
}
