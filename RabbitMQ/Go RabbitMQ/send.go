package main

import "log"

import amqp "github.com/rabbitmq/amqp091-go"

// 辅助函数来检查每个amqp调用的返回值：
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// 1. 尝试连接RabbitMQ，建立连接
// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
conn, err := amqp.Dial("amqp://guest:guest@47.93.20.204:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()

// 2. 接下来，我们创建一个通道，大多数API都是用过该通道操作的。
ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()

// 3. 声明消息要发送到到队列
q, err := c

func main() {

}