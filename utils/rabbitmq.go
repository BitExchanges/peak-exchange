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
		Default map[string]string `yaml:"default"`
		Kline   map[string]string `yaml:"kline"`
		Ticker  map[string]string `yaml:"ticker"`
	} `yaml:"exchange"`

	Queue struct {
		Kline  map[string]string `yaml:"kline"`
		Ticker map[string]string `yaml:"ticker"`
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
		return fmt.Errorf("发送队列消息失败: %s", err)
	}
	return nil
}

// declare RabbitMQ exchange
func DeclareExchange() error {
	channel, err := RabbitMqConnect.Channel()
	if err != nil {
		return fmt.Errorf("获取队列通道失败: %s", err)
	}

	err = channel.ExchangeDeclare(AmqpGlobalConfig.Exchange.Default["key"],
		AmqpGlobalConfig.Exchange.Default["type"],
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return fmt.Errorf("创建交换机 [default] 失败: %s", err)
	}
	log.Printf("创建交换机 [%s] 成功\n", AmqpGlobalConfig.Exchange.Default["key"])

	err = channel.ExchangeDeclare(AmqpGlobalConfig.Exchange.Kline["key"],
		AmqpGlobalConfig.Exchange.Kline["type"],
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return fmt.Errorf("创建交换机 [kline] 失败: %s", err)
	}

	log.Printf("创建交换机 [%s] 成功\n", AmqpGlobalConfig.Exchange.Kline["key"])

	err = channel.ExchangeDeclare(AmqpGlobalConfig.Exchange.Ticker["key"],
		AmqpGlobalConfig.Exchange.Ticker["type"],
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return fmt.Errorf("创建交换机 [ticker] 失败: %s", err)
	}
	log.Printf("创建交换机 [%s] 成功\n", AmqpGlobalConfig.Exchange.Ticker["key"])
	DeclareQueue(channel)
	return nil
}

func DeclareQueue(channel *amqp.Channel) error {
	var err error
	_, err = channel.QueueDeclare(AmqpGlobalConfig.Queue.Kline["key"], true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("创建队列 [kline] 失败: %s", err)
	}

	log.Printf("创建队列 [%s] 成功\n", AmqpGlobalConfig.Queue.Kline["key"])

	_, err = channel.QueueDeclare(AmqpGlobalConfig.Queue.Ticker["key"], true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("创建队列 [ticker] 失败: %s", err)
	}
	log.Printf("创建队列 [%s] 成功\n", AmqpGlobalConfig.Queue.Ticker["key"])
	BindQueue(channel)
	return nil
}

//  bind queue with exchange by routeKey
func BindQueue(channel *amqp.Channel) {
	channel.QueueBind(AmqpGlobalConfig.Queue.Kline["key"], "kline", AmqpGlobalConfig.Exchange.Kline["key"], false, nil)
	channel.QueueBind(AmqpGlobalConfig.Queue.Ticker["key"], "ticker", AmqpGlobalConfig.Exchange.Ticker["key"], false, nil)
}
